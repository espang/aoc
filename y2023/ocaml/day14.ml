open Core

let range i = List.init i ~f:(fun x -> x)
let range_rev i = List.init i ~f:(fun x -> i - x -1)

type cell = RoundRock | Rock | Empty
type dir = North | West | South | East
let cell_of_char = function 
  | '.' -> Empty | '#' -> Rock | 'O' -> RoundRock
  | _ -> failwith "unexpected char"
let char_of_cell = function | Empty -> '.' | Rock -> '#' | RoundRock -> 'O'

module Board = struct
  type t =
    {matrix: cell array array
    ; nrows: int
    ; ncols: int}

  let at t row col = try t.matrix.(row).(col) with Invalid_argument _ -> Rock
  let swap t (row, col) (row2, col2) =
    let tmp = t.matrix.(row).(col) in
    t.matrix.(row).(col) <- t.matrix.(row2).(col2);
    t.matrix.(row2).(col2) <- tmp
  let parse s =
    let rows =
      String.split_lines s
      |> List.map ~f:String.to_list
      |> List.map ~f:(List.map ~f:cell_of_char)
    in
    let nrows = List.length (List.hd_exn rows) in
    let ncols = List.length rows in
    let arr = Array.make_matrix ~dimx:nrows ~dimy:ncols Empty in
    List.cartesian_product (range ncols) (range nrows)
    |> List.iter ~f:(fun (col, row) ->
      arr.(row).(col) <- (List.nth_exn (List.nth_exn rows row) col));
    {matrix=arr;nrows=nrows;ncols=ncols}

  let is_round t (row, col) = phys_equal t.matrix.(row).(col) RoundRock
  let is_empty t row col = phys_equal t.matrix.(row).(col) Empty

  let show t =
    Array.iter t.matrix ~f:(fun row ->
      Array.iter row ~f:(fun c -> Printf.printf "%c" (char_of_cell c));
      Printf.printf "\n")

  let score t =
    Array.foldi t.matrix ~init:0 ~f:(fun rowi acc row->
      acc +
      Array.fold row ~init:0 ~f:(fun acc v -> match v with
        | RoundRock -> acc + t.nrows - rowi
        | _ -> acc))
   
  let tilt dir t =
    let col_range =
      match dir with
      | West | North | South -> (range t.ncols)
      | East  -> (range_rev t.ncols)
    in
    let row_range =
      match dir with
      | East | West | North -> (range t.nrows)
      | South -> (range_rev t.nrows)
    in
    let (dr, dc) = 
      match dir with
      | North -> (-1, 0)
      | South -> (1, 0)
      | West  -> (0, -1)
      | East  -> (0, 1)
    in
    let rec bubble row col=
      let row2 = row + dr and col2 = col + dc in
      match at t row col, at t row2 col2 with
      | RoundRock, Empty ->
        (swap t (row, col) (row2, col2); bubble row2 col2)
      | _, _ -> ()
    in
    List.iter row_range ~f:(fun row ->
      List.iter col_range ~f:(fun col ->
        bubble row col));
    t

  let equal t1 t2 = 
    Array.for_alli t1.matrix ~f:(fun rowi row ->
      Array.for_alli row ~f:(fun coli v ->
        phys_equal v t2.matrix.(rowi).(coli)))
  
  let copy t =
    { matrix= Array.init t.nrows ~f:(fun row -> Array.copy t.matrix.(row))
    ; nrows=t.nrows
    ; ncols=t.ncols}
end

let cycle board = 
  board
  |> Board.tilt North
  |> Board.tilt West
  |> Board.tilt South
  |> Board.tilt East

let find_cycle (board_factory: unit -> Board.t) =
  let rec first tortoise hare =
    if Board.equal tortoise hare 
    then hare
    else first (cycle tortoise) (cycle (cycle hare))
  in
  let rec second mu tortoise hare =
    if Board.equal tortoise hare
    then (mu, tortoise) 
    else second (succ mu) (cycle tortoise) (cycle hare)
  in
  let rec third tortoise lam hare =
    if Board.equal tortoise hare
    then lam
    else third tortoise (succ lam) (cycle hare)
  in
  let hare2 = first (board_factory()) (cycle (board_factory())) in
  let (mu, tortoise2) = second 0 (board_factory()) (cycle hare2) in
  let lam = third tortoise2 1 (cycle (Board.copy tortoise2)) in
  (mu, lam)

let solve_part2 board (mu, lam) =
  let rec times n =
    if n = 0
    then ()
    else
      (let _ = cycle board in times (pred n))
  in
  times (mu + ((1000000000 - mu) % lam));
  board

let part1 content = 
  Board.parse content
  |> Board.tilt North
  |> Board.score
  |> Printf.printf "%d\n"

let part2 content =
  let fac = fun () -> Board.parse content in
  (solve_part2 (Board.parse content) (find_cycle fac))
  |> Board.score
  |> Printf.printf "%d\n"
