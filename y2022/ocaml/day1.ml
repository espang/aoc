
let day1 ?(top_n = 1) input=
  let foldf list s =
    match list with
    | [] -> [int_of_string s]
    | hd::tl ->
      match s with
      | "" -> 0 :: list
      | _ -> (hd + (int_of_string s)) :: tl
    in
  String.split_on_char '\n' input
  |> List.fold_left foldf []
  |> List.sort Int.compare
  |> List.rev
  |> List.to_seq
  |> Seq.take top_n
  |> Seq.fold_left (+) 0
  |> print_int

let part1 = day1 ~top_n:1
let part2 = day1 ~top_n:3
