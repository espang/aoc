open Core

type direction = North | South | West | East

type pipe =
  | Pipe of direction * direction
  | Ground
  | Start
  | In
  | Out

let pipe_to_char = function 
  | Pipe(North, South) | Pipe(South, North) -> '|'
  | Pipe(North, East)  | Pipe(East, North)  -> 'L'
  | Pipe(North, West)  | Pipe(West, North)  -> 'J'
  | Pipe(West, East)   | Pipe(East, West)   -> '-'
  | Pipe(West, South)  | Pipe(South, West)  -> '7'
  | Pipe(South, East)  | Pipe(East, South)  -> 'F'
  | Ground -> '.'      | Start -> 'S'
  | In -> 'I'          | Out -> 'O'
  | _ -> failwith "unexpected pipe"

let char_to_pipe = function 
  | '|' -> Pipe(North, South)  | 'L' -> Pipe(North, East) 
  | 'J' -> Pipe(North, West)   | '-' -> Pipe(West, East)  
  | '7' -> Pipe(West, South)   | 'F' -> Pipe(South, East)
  | '.' -> Ground      | 'S' -> Start 
  | _ -> failwith "unexecpted char"

let parse (content:string) = 
  let rows = String.split_lines content in
  let dimy = List.length rows in
  let dimx = String.length (List.hd_exn rows) in
  let arr = Array.make_matrix ~dimx ~dimy Ground in
  List.iteri rows ~f:(fun y row ->
    String.iteri row ~f:(fun x c ->
      arr.(x).(y) <- char_to_pipe c));
  arr

let show board =
  let dimx = Array.length board in
  let dimy = Array.length board.(0) in
  for y = 0 to dimy-1 do
    for x = 0 to dimx-1 do
      Printf.printf "%c" (pipe_to_char board.(x).(y))
    done;
    Printf.printf "\n"
  done

let find_start_exn board =
  Array.find_mapi_exn board 
    ~f:(fun x row ->
      match Array.findi row ~f:(fun _ v -> phys_equal v Start) with
      | Some (y, _) -> Some (x, y)
      | None -> None)

let walk_loop (x, y, direction) board =
  let move_to x y = function
    | North -> (x, y-1, South)
    | South -> (x, y+1, North)
    | West  -> (x-1, y, East)
    | East  -> (x+1, y, West)
  in
  let direction_to d x y =
    try 
      match board.(x).(y) with
      | Pipe(a, b) when (phys_equal a d) -> Some b
      | Pipe(b, a) when (phys_equal a d) -> Some b
      | _ -> None
    with
      Invalid_argument _ -> None
  in
  let rec aux (x', y', direction) path =
    match direction_to direction x' y' with
    | Some direction' ->
      if x = x' && y = y' && List.length path > 0
      then Some path
      else aux (move_to x' y' direction') ((x', y')::path)
    | None -> None
  in
  aux (x, y, direction) []

let loop board = 
  let (start_x, start_y) = find_start_exn board in
  let possible_starts = [
    (North, South);
    (South, North);
    (North, East);
    (East, North);
    (North, West);
    (West, North);
    (West, East);
    (East, West);
    (West, South);
    (South, West);
    (South, East);
    (East, South)
  ] in
  let solve_potential_board (from, towards) =
    board.(start_x).(start_y) <- Pipe(from, towards);
    match walk_loop (start_x, start_y, from) board with
    | Some path -> Some (List.length path / 2)
    | None -> None
  in
  List.find_map possible_starts ~f:solve_potential_board

let part1 content = 
  parse content
  |> loop
  |> Option.value_exn
  |> Printf.printf "%d\n"

type cell = Blocked | Free

let cell_to_char = function
  | Blocked -> '#'
  | Free    -> '.'

let transform board path =
  let dimx = 3 * Array.length board in
  let dimy = 3 * Array.length board.(0) in
  let board_3x3 = Array.make_matrix ~dimx ~dimy Free in
  Array.iteri board ~f:(fun x row ->
    Array.iteri row ~f:(fun y cell ->
      if List.exists path ~f:(fun (x2, y2) -> x = x2 && y = y2)
      then
        match cell with
        | Pipe(North, South) ->
          board_3x3.(x*3+1).(y*3) <- Blocked;
          board_3x3.(x*3+1).(y*3+1)   <- Blocked;
          board_3x3.(x*3+1).(y*3+2) <- Blocked
        | Pipe(North, East) ->
          board_3x3.(x*3+1).(y*3) <- Blocked;
          board_3x3.(x*3+1).(y*3+1)   <- Blocked;
          board_3x3.(x*3+2).(y*3+1) <- Blocked
        | Pipe(North, West) ->
          board_3x3.(x*3+1).(y*3) <- Blocked;
          board_3x3.(x*3+1).(y*3+1)   <- Blocked;
          board_3x3.(x*3).(y*3+1) <- Blocked
        | Pipe(West, East) ->
          board_3x3.(x*3).(y*3+1) <- Blocked;
          board_3x3.(x*3+1).(y*3+1)   <- Blocked;
          board_3x3.(x*3+2).(y*3+1) <- Blocked
        | Pipe(West, South) ->
          board_3x3.(x*3).(y*3+1) <- Blocked;
          board_3x3.(x*3+1).(y*3+1)   <- Blocked;
          board_3x3.(x*3+1).(y*3+2) <- Blocked
        | Pipe(South, East) ->
          board_3x3.(x*3+2).(y*3+1) <- Blocked;
          board_3x3.(x*3+1).(y*3+1)   <- Blocked;
          board_3x3.(x*3+1).(y*3+2) <- Blocked
        | _ -> failwith "unexpected field" 
      else ()));
  board_3x3

let show_3x3 board =
  let dimx = Array.length board in
  let dimy = Array.length board.(0) in
  for y = 0 to dimy-1 do
    for x = 0 to dimx-1 do
      Printf.printf "%c" (cell_to_char board.(x).(y))
    done;
    Printf.printf "\n"
  done

let solve_2 board =
  let (start_x, start_y) = find_start_exn board in
  let neighbours x y = [
    (x-1, y);
    (x, y-1);
    (x+1, y);
    (x, y+1)
  ] in
  (* this step also replaces the start with the appropriate pipe *)
  let path_length = Option.value_exn (loop board) * 2 in
  let direction = match board.(start_x).(start_y) with
    | Pipe (_, d) -> d
    | _ -> failwith "expect pipe at start!"
  in
  let path = Option.value_exn (walk_loop (start_x, start_y, direction) board) in
  let board_3x3 = transform board path in
  let is_free x y = 
    try
      phys_equal board_3x3.(x).(y) Free
    with
      Invalid_argument _ -> false
  in
  let set_blocked x y = 
    try
      board_3x3.(x).(y) <- Blocked
    with
      Invalid_argument _ -> ()
  in
  (* 0,0 will always be free and is connected to all fields outside of the loop *)
  let queue = Queue.create () in
  Queue.enqueue queue (0, 0);
  let rec aux () =
    if Queue.is_empty queue
    then ()
    else 
      let (x, y) = Queue.dequeue_exn queue in
      set_blocked x y;
      neighbours x y
      |> List.iter ~f:(fun (x2, y2) ->
        if is_free x2 y2
        then 
          (set_blocked x2 y2;
          Queue.enqueue queue (x2, y2))
        else ());
      aux ()
  in
  aux ();
  let transform_coordinate x y = (x / 3, y / 3) in
  (Array.foldi board_3x3 ~init:[] ~f:(fun x lst row ->
    Array.foldi row ~init:[] ~f:(fun y lst' cell ->
      match cell with
      | Free -> (transform_coordinate x y) :: lst'
      | Blocked -> lst')
    |> List.append lst)
  |> List.dedup_and_sort ~compare:(fun (x, y) (x2, y2) ->
      match compare x x2 with
      | 0 -> compare y y2
      | v -> v)
  |> List.length) - path_length
  
let part2 content = 
  parse content
  |> solve_2
  |> Printf.printf "%d\n"