open Core

type mapping =
  { lower_bound: int;
    upper_bound: int;
    delta: int;
  }

let mapping_contains value m =
  value >= m.lower_bound && value <= m.upper_bound

let numbers_from_string s =
  String.split s ~on:' '
  |> List.map ~f:int_of_string

let parse_mapping_forward s = 
  match numbers_from_string s with
  | [dest; src; len] -> 
    {lower_bound=src;
     upper_bound=(src + len - 1);
     delta=(dest - src)}
  | _ -> failwith "not a triple"

let parse_mappings f s=
  String.split_lines s
  |> List.tl_exn
  |> List.map ~f

let parse_seeds s = 
  String.chop_prefix_exn s ~prefix:"seeds: "
  |> numbers_from_string

let parse_content s f= 
  let parts = Str.split (Str.regexp "\n\n") s in
  (parse_seeds (List.hd_exn parts), 
   List.map (List.tl_exn parts) ~f:(parse_mappings f))

let map' value mappings =
  match List.find mappings ~f:(mapping_contains value) with
  | Some m -> value + m.delta
  | None -> value

let map_all lst_of_mappings value =
  List.fold_left lst_of_mappings ~f:map' ~init:value

let part1 content =
  let (seeds, lst_of_mappings) = parse_content content parse_mapping_forward in
  List.map seeds ~f:(map_all lst_of_mappings)
  |> List.min_elt ~compare:Int.compare
  |> Option.value_exn
  |> Printf.printf "%d\n"

let parse_mapping_backward s = 
  match numbers_from_string s with
  | [dest; src; len] -> 
    {lower_bound=dest;
      upper_bound=(dest + len - 1);
      delta=(src - dest)}
  | _ -> failwith "not a triple"

let convert_seeds seeds = 
  let rec convert_seeds' tuples = function
    | [] -> tuples
    | _ :: [] -> failwith "unexpected seed list"
    | start :: len :: tl ->
      convert_seeds' ((start, (start+len-1))::tuples) tl
  in
  convert_seeds' [] seeds

let part2 content = 
  let (seeds, lst_of_mappings) = parse_content content parse_mapping_backward in
  let seeds' = convert_seeds seeds in
  let in_range v (lower_bound, upper_bound) =
    lower_bound <= v && v <= upper_bound
  in
  let is_valid_seed v =
    List.exists seeds' ~f:(in_range v)
  in
  let lst_of_mappings' = List.rev lst_of_mappings
  in
  let handle_mappings v mappings =
    match List.find mappings ~f:(mapping_contains v) with
    | Some m -> v + m.delta
    | None -> v
  in
  let rec find_smallest_value i =
    match List.fold_left lst_of_mappings' ~init:i ~f:handle_mappings |> is_valid_seed with
    | true  -> i
    | false -> find_smallest_value (i + 1)
  in
  find_smallest_value 0
  |> Printf.printf "%d\n"