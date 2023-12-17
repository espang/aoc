open Core

type cell =
  | Empty
  | SplitterHorizontal
  | SplitterVertical
  | MirrorLeftRight (* \ *)
  | MirrorRightLeft (* / *)

type direction = North | South | West | East  

type position = {row:int; col:int}

let int_of_direction = function
  | North -> 1
  | South -> 2
  | West  -> 3
  | East  -> 4

let compare_pos pos1 pos2 =
  match compare pos1.row pos2.row with
  | 0 -> compare pos1.col pos2.col
  | n -> n

let move_dir positon = function
  | North -> { positon with row = positon.row - 1 }
  | South -> { positon with row = positon.row + 1 }
  | West  -> { positon with col = positon.col - 1 }
  | East  -> { positon with col = positon.col + 1 }

let mirror_dir cell direction =
  match cell, direction with
  (* / *)
  | MirrorLeftRight, North -> West
  | MirrorLeftRight, South -> East
  | MirrorLeftRight, West  -> North
  | MirrorLeftRight, East  -> South
  (* \ *)
  | MirrorRightLeft, North -> East
  | MirrorRightLeft, South -> West
  | MirrorRightLeft, West  -> South
  | MirrorRightLeft, East  -> North
  | _ -> direction

let transform c =
  match c with
  | '.'  -> Empty
  | '|'  -> SplitterVertical
  | '-'  -> SplitterHorizontal
  | '\\' -> MirrorLeftRight
  | '/'  -> MirrorRightLeft
  | _ -> failwith "unexpected char" 

let move_through dir pos board =
  try
    let cell_value = board.(pos.row).(pos.col) in
    match cell_value with
    | Empty -> [ (dir, move_dir pos dir) ]
    | SplitterHorizontal -> (match dir with
      | North | South -> [ (East, move_dir pos East); (West, move_dir pos West) ]
      | East  | West  -> [ (dir, move_dir pos dir) ])
    | SplitterVertical -> (match dir with
      | East  | West  -> [ (North, move_dir pos North); (South, move_dir pos South) ]
      | North | South -> [ (dir, move_dir pos dir) ])
    | MirrorLeftRight
    | MirrorRightLeft ->
      let new_dir = mirror_dir cell_value dir in
      [ (new_dir, move_dir pos new_dir)]
  with Invalid_argument _ -> []

let equal (dir1, pos1) (dir2, pos2) =
  phys_equal dir1 dir2 &&
  pos1.row = pos2.row &&
  pos1.col = pos2.col

let solve ?(start = (East, {row=0; col=0})) board=
  let on_board {row;col} =
    try let _ = board.(row).(col) in true with Invalid_argument _ -> false
  in
  let queue = Queue.create () in
  Queue.enqueue queue start;
  let rec aux energized =
    if Queue.is_empty queue
    then energized
    else
      let (dir, pos) = Queue.dequeue_exn queue in
      if on_board pos
      then
        (move_through dir pos board
        |> List.filter ~f:(fun (dir, {row; col}) ->
            not (Set.mem energized (int_of_direction dir, row, col)))
        |> Queue.enqueue_all queue;
        aux (Set.add energized (int_of_direction dir, pos.row, pos.col)))
      else
        aux energized
  in
  let (dir, {row; col}) = start in
  let s = Set.add (Set.empty (module Aoc.Triple.Triple)) (int_of_direction dir, row, col) in
  aux s
  |> Set.map (module Aoc.Pair.Pair) ~f:(fun (_, row, pos) -> (row, pos))
  |> Set.length

let part1 content = 
  Aoc.Parsing.parse_string_matrix transform content
  |> solve
  |> Printf.printf "%d\n"

let solve2 board =
  let shape = Aoc.Parsing.shape_of board in
  [
    List.map (Aoc.Parsing.range shape.nrows) ~f:(fun row ->
      (East, {row=row;col=0}));
    List.map (Aoc.Parsing.range shape.nrows) ~f:(fun row ->
      (West, {row=row;col=shape.ncols-1}));
    List.map (Aoc.Parsing.range shape.ncols) ~f:(fun col ->
      (South, {row=0;col=col}));
    List.map (Aoc.Parsing.range shape.ncols) ~f:(fun col ->
      (North, {row=shape.nrows-1;col=col}));
  ]
  |> List.concat
  |> List.map ~f:(fun start -> solve ~start board)
  |> List.max_elt ~compare
  |> Option.value_exn

let part2 content =
  Aoc.Parsing.parse_string_matrix transform content
  |> solve2
  |> Printf.printf "%d\n"
