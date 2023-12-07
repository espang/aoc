open Core

type handType =
  | FiveOfAKind
  | FourOfAKind
  | FullHouse
  | ThreeOfAKind
  | TwoPair
  | OnePair
  | HighCard

let handType_to_number = function
  | FiveOfAKind -> 1
  | FourOfAKind -> 2
  | FullHouse -> 3
  | ThreeOfAKind -> 4
  | TwoPair -> 5
  | OnePair -> 6
  | HighCard -> 7

let compare_hand_type ht1 ht2 =
  let n1 = handType_to_number ht1 in
  let n2 = handType_to_number ht2 in
  compare n1 n2

let card_to_number joker_value = function
  | 'A' -> 1
  | 'K' -> 2
  | 'Q' -> 3
  | 'J' -> joker_value
  | 'T' -> 5
  | '9' -> 6
  | '8' -> 7
  | '7' -> 8
  | '6' -> 9
  | '5' -> 10
  | '4' -> 11
  | '3' -> 12
  | '2' -> 13
  | _ -> failwith "unexpected card"

let compare_card card_to_number c1 c2 = 
  compare (card_to_number c1) (card_to_number c2)

let compare_cards card_to_number cards1 cards2 =
  let rec aux cs1 = function
    | [] -> if List.is_empty cs1 then 0 else 1
    | hd2::tl2 -> match cs1 with
      | [] -> -1
      | hd1::tl1 -> match compare_card card_to_number hd1 hd2 with
        | 0 -> aux tl1 tl2
        | v -> v
  in
  aux (String.to_list cards1) (String.to_list cards2)


type hand =
  { cards: string;
    handType: handType;
  }

let compare_hands card_to_number hand1 hand2 =
  match compare_hand_type hand1.handType hand2.handType with
  | 0 -> compare_cards card_to_number hand1.cards hand2.cards
  | v -> v
let frequencies hand =
  let update_frequency_counter = function
    | None -> 1
    | Some n -> n + 1
  in
  let m = Map.empty (module Char) in
  String.fold hand ~init:m ~f:(fun m' c -> 
    Map.update m' c ~f:update_frequency_counter)
  |> Map.to_alist
  |> List.map ~f:(fun (_, count) -> count)
  |> List.sort ~compare
  |> List.rev

let list_to_string lst =
  List.map lst ~f:string_of_int
  |> String.concat ~sep:", "

let make_hand frequencies s =
  match frequencies s with
  | [5] -> {cards=s;handType=FiveOfAKind}
  | [4;1] -> {cards=s;handType=FourOfAKind}
  | [3;2] -> {cards=s;handType=FullHouse}
  | [3;1;1] -> {cards=s;handType=ThreeOfAKind}
  | [2;2;1] -> {cards=s;handType=TwoPair}
  | [2;1;1;1] -> {cards=s;handType=OnePair}
  | [1;1;1;1;1] -> {cards=s;handType=HighCard}
  | v -> failwith (Printf.sprintf "unexpected hand: [%s] '%s'" (list_to_string v) s)

let must_split_into_two s on =
  match String.split s ~on with
  | [fst; snd] -> (fst, snd)
  | _ -> failwith "could not split into two"
  
let parse_line make_hand l =
  let (cards, bid) = must_split_into_two l ' ' in
  (make_hand cards, int_of_string bid)

let rank_hands card_to_number lst =
  let compare_fn (h1, _) (h2, _) =
    -1 * (compare_hands card_to_number h1 h2)
  in
  List.sort lst ~compare:compare_fn

let part1 content =
  String.split_lines content
  |> List.map ~f:(parse_line (make_hand frequencies))
  |> (rank_hands (card_to_number 4))
  |> List.foldi ~init:0 ~f:(fun rank acc (_, bid) -> acc + (rank + 1) * bid)
  |> Printf.printf "%d\n"

let frequencies2 hand =
  let update_frequency_counter = function
    | None -> 1
    | Some n -> n + 1
  in
  let jokers = String.count hand ~f:(phys_equal 'J') in
  let add_jokers = function
    | hd::tl -> (hd + jokers) :: tl
    | [] -> [jokers]
  in
  let m = Map.empty (module Char) in
  String.filter hand ~f:(fun c -> not (phys_equal c 'J'))
  |> String.fold ~init:m ~f:(fun m' c -> 
    Map.update m' c ~f:update_frequency_counter)
  |> Map.to_alist
  |> List.map ~f:(fun (_, count) -> count)
  |> List.sort ~compare
  |> List.rev
  |> add_jokers

let part2 content = 
  String.split_lines content
  |> List.map ~f:(parse_line (make_hand frequencies2))
  |> (rank_hands (card_to_number 14))
  |> List.foldi ~init:0 ~f:(fun rank acc (_, bid) -> acc + (rank + 1) * bid)
  |> Printf.printf "%d\n"
