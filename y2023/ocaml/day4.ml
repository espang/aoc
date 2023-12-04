open Core

let must_split_into_two s on =
  match String.split s ~on with
  | [fst; snd] -> (fst, snd)
  | _ -> failwith "could not split into two"

let to_number_list s =
  String.split s ~on:' '
  |> List.filter ~f:(fun s -> not (String.is_empty s))
  |> List.map ~f:String.strip
  |> List.map ~f:int_of_string

let parse_line l =
  (* Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53 *)
  let (_, cards) = must_split_into_two l ':' in
  let (winning, numbers) = must_split_into_two cards '|' in
  (to_number_list winning, to_number_list numbers)

let matching_numbers (winning, numbers) = 
  let winning_set = Set.of_list (module Int) winning in
  let numbers_set = Set.of_list (module Int) numbers in
  Set.inter winning_set numbers_set
  |> Set.length

let score_cards cards =
  let i = matching_numbers cards in
  if i = 0 
  then 0.
  else 2. ** (float_of_int (i -1))

let part1 content = 
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:score_cards
  |> List.fold ~init:0. ~f:(+.)
  |> Printf.printf "%f" 

let do_part2 game_points =
  let arr = Array.create ~len:(List.length game_points) 1 in
  let update_card (idx, score) =
    let multiplier = arr.(idx) in
    try
      for i = 1 to score do
        arr.(idx+i) <- arr.(idx+i) + multiplier
      done
    with Invalid_argument _ -> ()
  in
  game_points
  |> List.iter ~f:update_card;
  arr

let part2 content =
  String.split_lines content
  |> List.map ~f:parse_line
  |> List.map ~f:matching_numbers
  |> List.mapi ~f:(fun i score -> (i, score))
  |> do_part2
  |> Array.fold ~init:0 ~f:(+)
  |> Printf.printf "%d" 
