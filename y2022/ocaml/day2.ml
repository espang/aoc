
type play = Rock | Paper | Scissors

let from_char = function
  | 'A' -> Rock
  | 'B' -> Paper
  | 'C' -> Scissors
  | _ -> failwith "unexpected char"

let winning_play = function
  | Rock -> Paper
  | Paper -> Scissors
  | Scissors -> Rock

let losing_play = function 
  | Rock -> Scissors
  | Paper -> Rock
  | Scissors -> Paper

let parse_line1 s = 
  let transform c = Char.chr (Char.code c - 23) in
  (from_char s.[0], transform s.[2] |> from_char)

let points_for (theirs, ours) =
  let points_for_play = match ours with
    | Rock -> 1 | Paper -> 2 | Scissors -> 3
  in
  points_for_play +
  if ours = (winning_play theirs)
  then 6
  else
    if ours = theirs
    then 3
    else 0

let part1 input =
  String.split_on_char '\n' input
  |> List.map parse_line1
  |> List.map points_for
  |> List.fold_left (+) 0
  |> print_int

let parse_line2 s = 
  let transformf = function
    | 'X' -> losing_play
    | 'Y' -> (fun c -> c)
    | 'Z' -> winning_play
    | _ -> failwith "unexpected char"
  in
  let their_play = from_char s.[0]
  in
  (their_play, (transformf s.[2]) their_play)

let part2 input =
  String.split_on_char '\n' input
  |> List.map parse_line2
  |> List.map points_for
  |> List.fold_left (+) 0
  |> print_int
