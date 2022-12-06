let str_to_char_list s = List.of_seq (String.to_seq s)

let stack_tests = ["NZ"; "DCM"; "P"]
let stack = ["RQGPCF";
  "PCTW"; "CMPHB";
  "RPMSQTL"; "NGVZJHP";
  "JPD"; "RTJFZPGL";
  "JTPFCHLN"; "WCTHQZVG"]

let stack_from inputs =
  let stacks = Array.make (List.length inputs) [] in
  List.iteri (fun i s -> stacks.(i) <- str_to_char_list s) inputs;
  stacks

let parse_move s = 
  (* move 1 from 2 to 1 *)
  let elements = String.split_on_char ' ' s in
  (int_of_string (List.nth elements 1),
  int_of_string (List.nth elements 3)-1,
  int_of_string (List.nth elements 5)-1)

let rec handle_move stacks (move, from, target) =
  if move = 0
  then ()
  else
    match stacks.(from) with
    | [] -> failwith "empty stack"
    | hd :: tl ->
      stacks.(target) <- hd :: stacks.(target);
      stacks.(from) <- tl;
      handle_move stacks (move-1, from, target)

let handle_move2 stacks (move, from, target) =
  let source_list = stacks.(from) in
  let to_move = List.to_seq source_list |> Seq.take move |> List.of_seq in
  let rest = List.to_seq source_list |> Seq.drop move |> List.of_seq in
  stacks.(target) <- to_move @ stacks.(target);
  stacks.(from) <- rest

let day5 handlef stacks input =
  let stacks' = Array.copy stacks in
  String.split_on_char '\n' input
  |> List.map parse_move
  |> List.iter (handlef stacks');
  Array.iter (fun list -> print_char (List.hd list)) stacks';
  print_newline ()

let part1 = day5 handle_move (stack_from stack)
let part2 = day5 handle_move2 (stack_from stack)
