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

let number_of_solutions_for (condition_records, damaged_springs) =
  let handle_operational counter damaged_counts =
    if counter = 0
    then Some damaged_counts
    else
      match damaged_counts with
      | [] -> None
      | v::tl ->
        if v = counter
        then Some tl
        else None
  in
  let rec count_solutions counter damaged condition_records =
    if List.is_empty damaged && counter > 0 then 0 else
    if counter > 0 && List.hd_exn damaged < counter then 0 else
    match condition_records with
    | [] ->
      if counter > 0
      then if List.length damaged = 1 && List.hd_exn damaged = counter then 1 else 0
      else if List.is_empty damaged then 1 else 0
    | hd::tl ->
      match hd with
      | Damaged -> count_solutions (counter + 1) damaged tl
      | Operational -> 
        (match handle_operational counter damaged with
        | Some damaged' -> count_solutions 0 damaged' tl
        | None -> 0)
      | Unknown ->
        (count_solutions (counter + 1) damaged tl +
        match handle_operational counter damaged with
        | Some damaged' -> count_solutions 0 damaged' tl
        | None -> 0)
  in
  count_solutions 0 damaged_springs condition_records

let part1 content = 
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:number_of_solutions_for
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"

let repeat5 lst =
  List.append lst lst
  |> List.append lst
  |> List.append lst
  |> List.append lst

let repeat5_with_unknown lst = 
  List.join [lst; [Unknown]; lst;[Unknown]; lst;[Unknown]; lst;[Unknown]; lst]

let extend (condition_records, damaged_springs) =
  (repeat5_with_unknown condition_records, repeat5 damaged_springs)

let part2 content =
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:extend
  |> (fun coll -> List.take coll 10)
  |> List.map ~f:number_of_solutions_for
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"

(* 
(* ?#??.??..##?.???#.? 4,1,2,1,1,1 *)
(* 1-4.0-2. 2-3 .1-4.
   1-6.0-2. 2-3 .1-4.
   1-6.0-2. 2-3 .1-4.
   1-6.0-2. 2-3 .1-4.
   1-6.0-2. 2-3 .1-4.
*)

?#.#???#???.??.???? 2,2,1,1,2,1 *)