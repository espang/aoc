
let read_whole_file filename =
  let ch = open_in filename in
  let s = really_input_string ch (in_channel_length ch) in
  close_in ch;
  s

let usage () =
  print_newline ();
  print_string "Expects exactly one argument to select the day to solve. for example 'day1'.";
  print_newline ()

let solve day =
  let (filename, part1, part2) = match day with
    | "day3" -> 
      ("../inputs/input3.txt", Day3.part1, Day3.part2)
    | "day4" -> 
      ("../inputs/input4.txt", Day4.d4p1, Day4.d4p2)
    | "day5" -> 
     ("../inputs/input5.txt", Day5.d5p1, Day5.d5p2)
    | _ -> failwith "unexpected day"
  in
  let input = read_whole_file filename in
  let () = print_newline () in
  let () = print_string "Solution for " in
  let () = print_string day in
  let () = print_newline () in
  let () = print_string "part1: " in
  let () = (part1 input) in
  let () = print_newline () in
  let () = print_string "part2: " in
  let () = (part2 input) in
  print_newline ()

let () =
  let args = Sys.argv in
  match Array.length args with
  | 1 -> usage (); exit 1
  | 2 -> solve args.(1)
  | _ -> usage (); exit 1
  
  