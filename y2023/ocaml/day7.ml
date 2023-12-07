open Core

let must_split_into_two s on =
  match String.split s ~on with
  | [fst; snd] -> (fst, snd)
  | _ -> failwith "could not split into two"

type handType = FiveOfAKind | FourOfAKind | FullHouse | ThreeOfAKind | TwoPair | OnePair | HighCard

let handType_to_number = function
  | FiveOfAKind -> 7 | FourOfAKind -> 6 | FullHouse -> 5 | ThreeOfAKind -> 4
  | TwoPair     -> 3 | OnePair     -> 2 | HighCard  -> 1

let card_to_number joker_value = function
  | 'A' -> 13 | 'K' -> 12 | 'Q' -> 11 | 'T' -> 9 | '9' -> 8 | '8' -> 7
  | '7' -> 6  | '6' -> 5  | '5' -> 4  | '4' -> 3 | '3' -> 2 | '2' -> 1  
  | 'J' -> joker_value | _ -> failwith "unexpected card"

let highest_card joker_value cards1 cards2 =
  let card_to_number' = card_to_number joker_value in
  let compare_card c1 c2 =  compare (card_to_number' c1) (card_to_number' c2) in
  let compare_cards cards1 cards2 =
    let rec aux = function
      | [], [] -> 0
      | [], _  -> -1
      | _, []  -> 1
      | (v1::cs1), (v2::cs2) ->
        match compare_card v1 v2 with
        | 0 -> aux (cs1, cs2)
        | v -> v
    in
    aux ((String.to_list cards1), (String.to_list cards2))
  in
  compare_cards cards1 cards2

let compare_hands joker_value (cs1, ht1) (cs2, ht2) =
  let compare_hand_type ht1 ht2 = 
    compare (handType_to_number ht1) (handType_to_number ht2)
  in
  match compare_hand_type ht1 ht2 with
  | 0 -> highest_card joker_value cs1 cs2
  | v -> v

let frequencies hand =
  let update_frequency_counter = function
    | None -> 1
    | Some n -> n + 1
  in
  let m = Map.empty (module Char) in
  String.fold hand ~init:m ~f:(Map.update ~f:update_frequency_counter)
  |> Map.data
  |> List.sort ~compare:Int.descending

let hand_type_from_freqs = function 
  | [5] -> FiveOfAKind | [4;1] -> FourOfAKind | [3;2] -> FullHouse
  | [3;1;1] -> ThreeOfAKind | [2;2;1] -> TwoPair | [2;1;1;1] -> OnePair
  | [1;1;1;1;1] -> HighCard
  | _ -> failwith "unexpected hand"

let parse_line freqs l =
  let (cards, bid) = must_split_into_two l ' ' in
  ((cards, hand_type_from_freqs (freqs cards)), int_of_string bid)

let rank_hands joker_value lst =
  let compare_fn (h1, _) (h2, _) =
    compare_hands joker_value h1 h2
  in
  List.sort lst ~compare:compare_fn

let frequencies_with_jokers hand =
  let without_jokers = String.filter hand ~f:(fun c -> not (phys_equal c 'J')) in
  let jokers = (5 - String.length without_jokers) in
  match frequencies without_jokers with
  | hd::tl -> (hd + jokers) :: tl
  | [] -> [jokers]
  
let count_winnings content joker_value freq =
  String.split_lines content
  |> List.map ~f:(parse_line freq)
  |> (rank_hands joker_value)
  |> List.foldi ~init:0 ~f:(fun rank acc (_, bid) -> acc + (rank + 1) * bid)

let part1 content = Printf.printf "%d\n" (count_winnings content 10 frequencies)
let part2 content = Printf.printf "%d\n" (count_winnings content 0 frequencies_with_jokers)
