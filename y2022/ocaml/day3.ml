module CS = Set.Make(Char)

let priority c =
  let code = Char.code c in
  if code >= Char.code 'a' && code <= Char.code 'z'
  then code - (Char.code 'a') + 1
  else
    if code >= Char.code 'A' && code <= Char.code 'Z'
    then code - (Char.code 'A') + 27
    else failwith "unexpected item"

let str_to_char_set s = 
  List.fold_right CS.add (List.of_seq (String.to_seq s)) CS.empty 

let parseLine s =
  let len = String.length s in
  if len mod 2 = 0
  then (str_to_char_set (String.sub s 0 (len/2)),
    str_to_char_set  (String.sub s (len/2) (len/2)))
  else failwith "unexpected length of line"

let important_item (s1, s2) =
  let intersect = CS.inter s1 s2 in
  if CS.cardinal intersect = 1
  then List.hd (CS.elements intersect)
  else failwith "expected exactly one element"

let part1 input =
  let lines = String.split_on_char '\n' input in
  lines
  |> List.map parseLine
  |> List.map important_item
  |> List.map priority
  |> List.fold_left (+) 0
  |> print_int

let find_priority s1 s2 s3 =
  let intersect = CS.inter (str_to_char_set s1)
    (CS.inter (str_to_char_set s2) (str_to_char_set s3))
  in
  if CS.cardinal intersect = 1
  then priority (List.hd (CS.elements intersect))
  else failwith "expected exactly one element"
  
let priority_of_groups list =
  let rec ig list acc =
    match list with
    | a::b::c::tl ->
      ig tl ( (find_priority a b c) :: acc)
    | [] -> acc
    | _ -> failwith "x"
    in
  ig list []
  
let part2 input =
  let lines = String.split_on_char '\n' input in
  lines
  |> priority_of_groups
  |> List.fold_left (+) 0
  |> print_int