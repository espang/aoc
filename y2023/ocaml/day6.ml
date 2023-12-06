open Core

type race =
  { time: int;
    distance: int;
  }
  let number_of_ways_to_win {time; distance} =
    let hold_time_to_distance ht =
      ht * (time - ht)
    in
    let total_winnings fst =
      (* the distances we can reach is symmetric
         so if the first winning option is on index
         n from the start is also on index -n form the
        end.*)
      let number_of_options = time + 1 in
      number_of_options - 2 * fst
    in
    match Seq.init time (fun x -> x)
      |> Seq.find (fun t -> (hold_time_to_distance t) > distance) with
    | Some first_win -> total_winnings first_win
    | None -> 0
  
let part1 _ =
  [
    {time=54; distance=302};
    {time=94; distance=1476};
    {time=65; distance=1029};
    {time=92; distance=1404}
  ] 
  |> List.map ~f:number_of_ways_to_win
  |> List.fold ~init:1 ~f:( * )
  |> Printf.printf "%d\n"

let part2 _ =
  {time=54946592; distance=302147610291404}
  |> number_of_ways_to_win
  |> Printf.printf "%d\n"

