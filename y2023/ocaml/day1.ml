let conversion s =
  match s with
  | "one" -> "1"   | "two" -> "2"   | "three" -> "3"
  | "four" -> "4"  | "five" -> "5"  | "six" -> "6"
  | "seven" -> "7" | "eight" -> "8" | "nine" -> "9"
  | _ -> s

let first_digit r l =
  let _ = Str.search_forward r l 0 in
  let found = Str.matched_string l in
  int_of_string (conversion found)

let last_digit r l =
  let _ = Str.search_backward r l (String.length l) in
  let found = Str.matched_string l in
  int_of_string (conversion found)

let handle_line r l =
  let first = first_digit r l in
  let last = last_digit r l in
  first * 10 + last

let part1 content = 
  let r = Str.regexp {|\([0-9]\)|} in
  String.split_on_char '\n' content
  |> List.map (handle_line r)
  |> List.fold_left (fun acc v -> acc + v) 0
  |> Printf.printf "%d"

let part2 content =
  let r = Str.regexp {|\(one\|two\|three\|four\|five\|six\|seven\|eight\|nine\|[0-9]\)|} in
  String.split_on_char '\n' content
  |> List.map (handle_line r)
  |> List.fold_left (fun acc v -> acc + v) 0
  |> Printf.printf "%d"
