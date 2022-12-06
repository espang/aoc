module CS = Set.Make(Char)

type tracker = 
  { elements: CS.t
  ; ordered: char Queue.t
  ; index: int}

let make_tracker () = {
  elements = CS.empty;
  ordered = Queue.create ();
  index = 0;
}

let found n {elements; index; _} =
  if CS.cardinal elements = n
  then
    Some index 
  else None

let rec remove_until {elements; ordered; index} c =
  let c' = Queue.pop ordered in
  if Char.equal c c'
  then
    let () = Queue.add c ordered in
    {elements; ordered; index}
  else
    remove_until {elements = CS.remove c' elements;ordered;index} c

let add_element c {elements; ordered; index} =
  if CS.mem c elements
  then 
    remove_until {elements; ordered; index=Int.succ index} c
  else
    let () = Queue.add c ordered in
    {elements = CS.add c elements;
     ordered = ordered;
    index = Int.succ index}

let solve n input = 
  let rec s t = function
    | [] -> failwith "no solution"
    | hd :: tl ->
      let new_t = add_element hd t in
      match found n new_t with
      | Some i -> i
      | None -> s new_t tl
  in
  s (make_tracker ()) (List.of_seq (String.to_seq input))

let solve_part1 input = 
  let index = solve 4 input in
  print_int index

let solve_part2 input =
  let index = solve 14 input in
  print_int index
