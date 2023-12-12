open Core

let must_split_into_two s on =
  match String.split s ~on with
  | [fst; snd] -> (fst, snd)
  | _ -> failwith "could not split into two"

type spring = Operational | Damaged | Unknown

let spring_of_char = function
  | '.' -> Operational
  | '#' -> Damaged
  | '?' -> Unknown
  | _ -> failwith "unexpected char"

let parse_line l =
  let (raw_condition_records, raw_damaged_springs) = must_split_into_two l ' ' in
  let condition_records =
    String.to_list raw_condition_records
    |> List.map ~f:spring_of_char
  in
  let damaged_springs =
    String.split raw_damaged_springs ~on:','
    |> List.map ~f:int_of_string
  in
  (condition_records, damaged_springs)

let repeat5 lst =
  List.append lst lst
  |> List.append lst
  |> List.append lst
  |> List.append lst

let repeat5_with_unknown lst = 
  List.join [
    lst; [Unknown];
    lst; [Unknown];
    lst; [Unknown];
    lst; [Unknown];
    lst
  ]

let extend (condition_records, damaged_springs) =
  (repeat5_with_unknown condition_records, repeat5 damaged_springs)

module Pair = struct
  module T = struct
    type t = spring list * int list

    let rec list_equal eq coll1 coll2 =
      match coll1, coll2 with
      | [], [] -> true
      | _,  [] -> false
      | [], _  -> false
      | (v1::cs1), (v2::cs2) ->
        if eq v1 v2
        then list_equal eq cs1 cs2
        else false
    let compare (cs1, ds1) (cs2, ds2) =
      if list_equal (=) ds1 ds2 && list_equal phys_equal cs1 cs2
      then 0
      else 1
    let sexp_of_t = 
      Tuple2.sexp_of_t (List.sexp_of_t sexp_of_opaque)  (List.sexp_of_t Int.sexp_of_t)
    
    let t_of_sexp = Tuple2.t_of_sexp (List.t_of_sexp opaque_of_sexp) (List.t_of_sexp Int.t_of_sexp)
  end

  include T
  include Comparable.Make(T)
end

let solve (condition_records, damaged_springs) =
  let is_operational v = phys_equal v Operational in
  let possible number conditions =
    if List.length conditions >= number
    then 
      match List.nth conditions number with
      | Some Damaged -> false
      | _ -> List.take conditions number |> List.exists ~f:is_operational |> not
    else false
  in
  let memory = ref (Map.empty (module Pair)) in
  let rec f = function 
    | ([], [])   -> 1
    | (cs, []) -> if List.exists cs ~f:(phys_equal Damaged) then 0 else 1
    | ([], _)  -> 0
    | (conds, dmgs) as t ->
      if Map.mem !memory t
      then Map.find_exn !memory t
      else
        let next_number = List.hd_exn dmgs in
        let data =
          match (List.hd_exn conds), possible next_number conds with
          | Operational, _
          | Unknown, false -> f (List.tl_exn conds, dmgs)
          | Damaged, true  -> f (List.drop conds (succ next_number), List.tl_exn dmgs)
          | Damaged, false -> 0
          | Unknown, true  -> 
            f (List.drop conds (succ next_number), List.tl_exn dmgs)
            + f (List.tl_exn conds, dmgs)
        in
        memory := Map.add_exn !memory ~key:t ~data;
        data
  in
  f (condition_records, damaged_springs)

let part1 content = 
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:solve
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"

let part2 content =
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:extend
  |> List.map ~f:solve
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"
