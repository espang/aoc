open Core

let value_from_char_list cl =
  Char.escaped (List.hd_exn cl) ^ Char.escaped (List.last_exn cl)
  |> int_of_string

let transform1 s =
  let rec transform result = function
    | []                    -> List.rev result
    | '0' .. '9' as v :: tl -> transform (v::result) tl
    | _ :: tl               -> transform result tl
  in
  transform [] (String.to_list s)
  |> value_from_char_list

let transform2 s =
  let rec transform result = function
    | []                                    -> List.rev result
    | '0' .. '9' as v :: tl                 -> transform (v::result) tl
    | 'o' :: 'n' :: 'e' :: tl               -> transform ('1'::result) ('e' :: tl)
    | 't' :: 'w' :: 'o' :: tl               -> transform ('2'::result) ('o' :: tl)
    | 't' :: 'h' :: 'r' :: 'e' :: 'e' :: tl -> transform ('3'::result) ('e' :: tl)
    | 'f' :: 'o' :: 'u' :: 'r' :: tl        -> transform ('4'::result) tl
    | 'f' :: 'i' :: 'v' :: 'e' :: tl        -> transform ('5'::result) ('e' :: tl)
    | 's' :: 'i' :: 'x' :: tl               -> transform ('6'::result) tl
    | 's' :: 'e' :: 'v' :: 'e' :: 'n' :: tl -> transform ('7'::result) tl
    | 'e' :: 'i' :: 'g' :: 'h' :: 't' :: tl -> transform ('8'::result) ('t' :: tl)
    | 'n' :: 'i' :: 'n' :: 'e' :: tl        -> transform ('9'::result) ('e' :: tl)
    | _ :: tl -> transform result tl
  in
  transform [] (String.to_list s)
  |> value_from_char_list

let apply content ft =
  String.split_lines content
  |> List.map ~f:ft
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d"

let part1 content = apply content transform1
let part2 content = apply content transform2