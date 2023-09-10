type range =
  { start:int
  ; upto:int}

let parse_range s = 
  match String.split_on_char '-' s with
  | [a;b] -> {start=int_of_string a; upto= int_of_string b}
  | _ -> failwith "invalid string for a range"

let parse_line s =
  match String.split_on_char ',' s with
  | [a;b] -> parse_range a, parse_range b
  | _ -> failwith "invalid string for a line"

let fully_contained (r1, r2 )=
  (r1.start <= r2.start && r1.upto >= r2.upto)
  || (r2.start <= r1.start && r2.upto >= r1.upto)

let any_overlap (r1, r2) =
  (r1.start >= r2.start && r1.start <= r2.upto)
  || (r2.start >= r1.start && r2.start <= r1.upto)

let day4 input select =
  String.split_on_char '\n' input
  |> List.map parse_line
  |> List.filter select
  |> List.length


let part1 input = print_int (day4 input fully_contained)
let part2 input = print_int (day4 input any_overlap)
