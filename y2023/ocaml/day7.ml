open Core

let must_split_into_two s on =
  match String.split s ~on with
  | [fst; snd] -> (fst, snd)
  | _ -> failwith "could not split into two"

type handType = FiveOfAKind | FourOfAKind | FullHouse | ThreeOfAKind | TwoPair | OnePair | HighCard

type hand =
  { cards: string;
    handType: handType;
  }

let handType_to_number = function
  | FiveOfAKind -> 7 | FourOfAKind -> 6 | FullHouse -> 5 | ThreeOfAKind -> 4
  | TwoPair     -> 3 | OnePair     -> 2 | HighCard  -> 1

let card_to_number joker_value = function
  | 'A' -> 13 | 'K' -> 12 | 'Q' -> 11 | 'J' -> joker_value
  | 'T' -> 9  | '9' -> 8  | '8' -> 7  | '7' -> 6
  | '6' -> 5  | '5' -> 4  | '4' -> 3  | '3' -> 2
  | '2' -> 1 | _ -> failwith "unexpected card"

let compare_hands card_to_number hand1 hand2 =
  let compare_cards card_to_number cards1 cards2 =
    let compare_card card_to_number c1 c2 = 
      compare (card_to_number c1) (card_to_number c2)
    in  
    let rec aux cs1 = function
      | [] -> if List.is_empty cs1 then 0 else 1
      | hd2::tl2 -> match cs1 with
        | [] -> -1
        | hd1::tl1 -> match compare_card card_to_number hd1 hd2 with
          | 0 -> aux tl1 tl2
          | v -> v
    in
    aux (String.to_list cards1) (String.to_list cards2)
  in
  let compare_hand_type ht1 ht2 = 
    compare (handType_to_number ht1) (handType_to_number ht2)
  in
  match compare_hand_type hand1.handType hand2.handType with
  | 0 -> compare_cards card_to_number hand1.cards hand2.cards
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

let make_hand frequencies s =
  match frequencies s with
  | [5]         -> {cards=s; handType=FiveOfAKind}
  | [4;1]       -> {cards=s; handType=FourOfAKind}
  | [3;2]       -> {cards=s; handType=FullHouse}
  | [3;1;1]     -> {cards=s; handType=ThreeOfAKind}
  | [2;2;1]     -> {cards=s; handType=TwoPair}
  | [2;1;1;1]   -> {cards=s; handType=OnePair}
  | [1;1;1;1;1] -> {cards=s; handType=HighCard}
  | _ -> failwith (Printf.sprintf "unexpected hand: '%s'" s)

let parse_line make_hand l =
  let (cards, bid) = must_split_into_two l ' ' in
  (make_hand cards, int_of_string bid)

let rank_hands card_to_number lst =
  let compare_fn (h1, _) (h2, _) =
    compare_hands card_to_number h1 h2
  in
  List.sort lst ~compare:compare_fn

let frequencies_with_jokers hand =
  let jokers = String.count hand ~f:(phys_equal 'J') in
  match frequencies (String.filter hand ~f:(fun c -> not (phys_equal c 'J'))) with
  | hd::tl -> (hd + jokers) :: tl
  | [] -> [jokers]
  
let count_winnings content joker_value freq =
  String.split_lines content
  |> List.map ~f:(parse_line (make_hand freq))
  |> (rank_hands (card_to_number joker_value))
  |> List.foldi ~init:0 ~f:(fun rank acc (_, bid) -> acc + (rank + 1) * bid)

let part1 content = Printf.printf "%d\n" (count_winnings content 10 frequencies)
let part2 content = Printf.printf "%d\n" (count_winnings content 0 frequencies_with_jokers)
