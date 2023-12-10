open Core

let test_input = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45"

let parse_line s = String.split s ~on:' ' |> List.map ~f:int_of_string

(* returns a list of the deltas between the integers in the list *)
let reduce_line lst =
  let rec aux result = function
    | [] -> List.rev result
    | _ :: [] -> List.rev result
    | fst::snd::tl -> aux ((snd - fst)::result) (snd::tl)
  in
  aux [] lst

let rec new_value lst =
  let is_zero lst = List.for_all lst ~f:(( = ) 0) in
  let lst' = reduce_line lst in
  if is_zero lst'
  then List.last_exn lst
  else
    let v' = List.last_exn lst in
    let v'' = new_value lst' in
    v' + v''

let part1 content = 
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:new_value
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"

let rec new_value_2 lst =
  let is_zero lst = List.for_all lst ~f:(( = ) 0) in
  let lst' = reduce_line lst in
  if is_zero lst'
  then List.hd_exn lst
  else
    let v' = List.hd_exn lst in
    let v'' = new_value_2 lst' in
    v' - v''

let part2 content = 
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:new_value_2
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"