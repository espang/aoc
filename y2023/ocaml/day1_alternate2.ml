open Core

let value_from_char_list cl =
  Char.escaped (List.hd_exn cl) ^ Char.escaped (List.last_exn cl)
  |> int_of_string

let is_digit = function | '0'..'9' -> true | _ -> false

let line_to_number s =
  String.filter s ~f:is_digit
  |> String.to_list
  |> value_from_char_list

let replacements =
    [("one", "o1e"); ("two", "t2o"); ("three", "t3e"); ("four", "4");
    ("five", "5e"); ("six", "6"); ("seven", "7"); ("eight", "e8t");
    ("nine", "9e")]

let transform line =
  let accumulate acc (pattern, with_) =
    String.substr_replace_all acc ~pattern ~with_
  in
  List.fold replacements ~init:line ~f:accumulate

let apply content ft =
  String.split_lines content
  |> List.map ~f:(Fn.compose line_to_number ft)
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d"

let part1 content = apply content Fn.id
let part2 content = apply content transform
