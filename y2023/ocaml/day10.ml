open Core
let range i = List.init i ~f:(fun x -> x)

type direction = North | South | West | East

type pipe =
  | Pipe of direction * direction
  | Ground
  | Start
  | In
  | Out

let char_of_pipe = function 
  | Pipe(North, South) | Pipe(South, North) -> '|'
  | Pipe(North, East)  | Pipe(East, North)  -> 'L'
  | Pipe(North, West)  | Pipe(West, North)  -> 'J'
  | Pipe(West, East)   | Pipe(East, West)   -> '-'
  | Pipe(West, South)  | Pipe(South, West)  -> '7'
  | Pipe(South, East)  | Pipe(East, South)  -> 'F'
  | Ground -> '.'      | Start -> 'S'
  | In -> 'I'          | Out -> 'O'
  | _ -> failwith "unexpected pipe"

let pipe_of_char = function 
  | '|' -> Pipe(North, South)  | 'L' -> Pipe(North, East) 
  | 'J' -> Pipe(North, West)   | '-' -> Pipe(West, East)  
  | '7' -> Pipe(West, South)   | 'F' -> Pipe(South, East)
  | '.' -> Ground      | 'S' -> Start 
  | _ -> failwith "unexecpted char"

module Board = struct
  type t =
    { matrix: pipe array array
    ; dimx: int
    ; dimy: int}
  let make ~dimx ~dimy =
    { matrix=Array.make_matrix ~dimx ~dimy Ground
    ; dimx=dimx
    ; dimy=dimy}
  let update t x y v = t.matrix.(x).(y) <- v
  let at t x y = try Some t.matrix.(x).(y) with Invalid_argument _ -> None
  let target_exn t x y dir =
    let v = t.matrix.(x).(y) in
    match v with 
    | Pipe(start, target) when (phys_equal dir start) -> target
    | Pipe(target, start) when (phys_equal dir start) -> target
    | _ -> raise (Invalid_argument "not a pipe")
  let start t =
    Array.find_mapi_exn t.matrix 
    ~f:(fun x row ->
      match Array.findi row ~f:(fun _ v -> phys_equal v Start) with
      | Some (y, _) -> Some (x, y)
      | None -> None)
  let path_loop t start_x start_y direction =
    let move x y = function
    | North -> (x, y-1, South)
    | South -> (x, y+1, North)
    | West  -> (x-1, y, East)
    | East  -> (x+1, y, West)
    in
    let rec aux (x, y, dir) path =
      try
        let dir' = target_exn t x y dir in
        if x = start_x && y = start_y && List.length path > 0
        then Some path
        else aux (move x y dir') ((x, y)::path)
      with
        Invalid_argument _ -> None
    in
    aux (start_x, start_y, direction) []
end

let parse (content:string) = 
  let rows = String.split_lines content in
  let dimy = List.length rows in
  let dimx = String.length (List.hd_exn rows) in
  let board = Board.make ~dimx ~dimy in
  List.iteri rows ~f:(fun y row ->
    String.iteri row ~f:(fun x c ->
      Board.update board x y (pipe_of_char c)));
  board

let find_start_and_path board = 
  let (start_x, start_y) = Board.start board in
  let possible_starts = [
    (North, South); (South, North); (North, East); (East, North);
    (North, West); (West, North); (West, East); (East, West);
    (West, South); (South, West); (South, East); (East, South)
  ] in
  let solve_potential_board (from, towards) =
    Board.update board start_x start_y (Pipe(from, towards));
    Board.path_loop board start_x start_y from
  in
  List.find_map_exn possible_starts ~f:solve_potential_board

let part1 content = 
  let path = find_start_and_path (parse content) in
  Printf.printf "%d\n" (List.length path / 2)

type cell = Blocked | Free
let cell_to_char = function
  | Blocked -> '#'
  | Free    -> '.'

let blocked_cells_of_pipe = function
  | Pipe(North, South) -> [(1, 0); (1, 1); (1, 2)]
  | Pipe(North, East)  -> [(1, 0); (1, 1); (2, 1)]
  | Pipe(North, West)  -> [(1, 0); (1, 1); (0, 1)]
  | Pipe(West, East)   -> [(0, 1); (1, 1); (2, 1)]
  | Pipe(West, South)  -> [(0, 1); (1, 1); (1, 2)]
  | Pipe(South, East)  -> [(1, 2); (1, 1); (2, 1)]
  | _ -> failwith "unexpected field"

module Board3x3 = struct
  type t =
  { matrix: cell array array
  ; dimx: int
  ; dimy: int}
  let make (board:Board.t) path =
    let dimx = 3 * board.dimx in
    let dimy = 3 * board.dimy in
    let board_3x3 = Array.make_matrix ~dimx ~dimy Free in
    List.cartesian_product (range board.dimx) (range board.dimy)
    |> List.iter ~f:(fun (x, y) ->
        if List.exists path ~f:(fun (x2, y2) -> x = x2 && y = y2)
        then
          let blocked_coords = blocked_cells_of_pipe board.matrix.(x).(y) in
          List.iter blocked_coords ~f:(fun (dx, dy) ->
            board_3x3.(x*3+dx).(y*3+dy) <- Blocked);
        else ()
      );
    { matrix=board_3x3
    ; dimx=3 * board.dimx
    ; dimy=3 * board.dimy}
  let on t x y = not (x < 0 || x >= t.dimx || y < 0 || y >= t.dimy)
  let is_free t x y =
    if on t x y
    then phys_equal t.matrix.(x).(y) Free
    else false
  let set_blocked t x y =
    if on t x y
    then t.matrix.(x).(y) <- Blocked
    else ()
  let free_coords t =
    List.cartesian_product (range t.dimx) (range t.dimy)
    |> List.filter ~f:(fun (x, y) -> is_free t x y)
  let mark_blocked t =
    let queue = Queue.create () in
    Queue.enqueue queue (0, 0);
    let neighbours x y = [(x-1, y);(x, y-1);(x+1, y);(x, y+1)] in
    let rec aux () =
      if Queue.is_empty queue
      then ()
      else 
        let (x, y) = Queue.dequeue_exn queue in
        set_blocked t x y;
        neighbours x y
        |> List.iter ~f:(fun (x2, y2) ->
          if is_free t x2 y2
          then 
            (set_blocked t x2 y2;
            Queue.enqueue queue (x2, y2))
          else ());
        aux ()
    in
    aux ()
end

let solve_2 board =
  let path = find_start_and_path board in
  let b3x3 = Board3x3.make board path in
  Board3x3.mark_blocked b3x3;
  let free_cords =
    (Board3x3.free_coords b3x3
    |> List.map ~f:(fun (x, y) -> (x / 3, y / 3))
    |> List.dedup_and_sort ~compare:(fun (x, y) (x2, y2) ->
        match compare x x2 with
        | 0 -> compare y y2
        | v -> v))
  in    
  (List.length free_cords) - (List.length path)

let part2 content = 
  parse content
  |> solve_2
  |> Printf.printf "%d\n"