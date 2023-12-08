open Core

type direction = Left | Right
type move = {left:string; right:string}
let parse_directions s = 
  let to_direction = function
    | 'L' -> Left
    | 'R' -> Right
    | _   -> failwith "unexpected direction"
  in
  String.to_list s
  |> List.map ~f:to_direction
  |> Sequence.cycle_list_exn

let parse_network s =
  let r = Str.regexp {|\(...\) = (\(...\), \(...\))|} in
  let parse_line line =
    if Str.string_match r line 0
    then
      (Str.matched_group 1 line,
        {left=Str.matched_group 2 line;
        right=Str.matched_group 3 line})
    else
      failwith "unexpected line"
  in
  String.split_lines s
  |> List.map ~f:parse_line
  |> Map.of_alist_exn (module String)

let parse_content s =
  match Str.split (Str.regexp "\n\n") s with
  | [instructions; network] ->
    (String.length instructions, parse_directions instructions, parse_network network)
  | _ -> failwith "unexpected content"

let cycle_length (l ,directions, network) target start =
  let move pos = function
    | Left -> (Map.find_exn network pos).left
    | Right -> (Map.find_exn network pos).right
  in
  let rec aux steps pos dirs =
    if (target pos) && (steps % l = 0)
    then steps
    else
      aux (steps + 1) (move pos (Sequence.hd_exn dirs)) (Sequence.tl_eagerly_exn dirs)
  in
  aux 0 start directions

let part1 content = 
  cycle_length (parse_content content) (String.equal "ZZZ") "AAA"
  |> Printf.printf "%d\n"

let lcm a b =
  let rec hcf a b =
    match a % b with
    | 0 -> b
    | rem -> hcf b rem
  in
  if a > b
  then a * b / (hcf a b)
  else a * b / (hcf b a)

let part2 content =
  let end_target s = phys_equal 'Z' (String.nget s 2) in
  let (_, _, network) as t = parse_content content in
  Map.keys network
  |> List.filter ~f:(fun s -> phys_equal 'A' (String.nget s 2))
  |> List.map ~f:(cycle_length t end_target)
  |> List.fold ~init:1 ~f:(fun acc v -> lcm acc v)
  |> Printf.printf "%d\n"
