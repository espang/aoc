open Core

let range i = List.init i ~f:(fun x -> x)

type cell = Ash | Rocks

let cell_of_char = function
  | '.' -> Ash
  | '#' -> Rocks
  | _ -> failwith "unexpected char"

module Board = struct
  type t = cell array array

  let make ~dimx ~dimy = Array.make_matrix ~dimx ~dimy Ash
  let shape t = (Array.length t, Array.length t.(0))

  let swap t x y = 
    t.(x).(y) <- match t.(x).(y) with
      | Ash -> Rocks
      | Rocks -> Ash

  let of_string s =
    let rows = 
      String.split_lines s
      |> List.map ~f:String.to_list
      |> List.map ~f:(List.map ~f:cell_of_char)
    in
    let dimy = List.length rows in
    let dimx = List.length (List.hd_exn rows) in
    let arr = make ~dimx ~dimy in
    List.cartesian_product (range dimx) (range dimy)
    |> List.iter ~f:(fun (col, row) ->
      arr.(col).(row) <- (List.nth_exn (List.nth_exn rows row) col));
    arr
end

let mirror_y board y =
  let (dimx, dimy) = Board.shape board in
  let on_board y' = y' >= 0 && y' < dimy in
  let rec aux dy =
    let y1 = y - dy in
    let y2 = y + 1 + dy in
    if on_board y1 && on_board y2
    then
      match List.exists (range dimx) ~f:(fun x -> not (phys_equal board.(x).(y1) board.(x).(y2))) with
      | true -> false
      | false -> aux (succ dy)
    else true
  in
  if y = dimy -1 then false else aux 0

let mirror_x board x =
  let (dimx, dimy) = Board.shape board in
  let on_board x' = x' >= 0 && x' < dimx in
  let rec aux dx =
    let x1 = x - dx in
    let x2 = x + 1 + dx in
    if on_board x1 && on_board x2
    then
      match List.exists (range dimy) ~f:(fun y -> not (phys_equal board.(x1).(y) board.(x2).(y))) with
      | true -> false
      | false -> aux (succ dx)
    else true
  in
  if x = dimx - 1 then false else aux 0

let value_of ?(skip_x = -1) ?(skip_y = -1) board =
  let (dimx, dimy) = Board.shape board in
  let skip to_skip f i = 
    if to_skip = i
    then false
    else f i
  in
  match List.find (range dimx) ~f:(skip skip_x (mirror_x board)) with
  | Some x -> (`x, x)
  | None -> 
    match List.find (range dimy) ~f:(skip skip_y (mirror_y board)) with
    | None -> failwith "no mirror line found"
    | Some y -> (`y, y)
  
let to_value = function
    | (`x, x) -> succ x
    | (`y, y) -> 100 * (succ y)

let part1 content =
  Str.split (Str.regexp "\n\n") content
  |> List.map ~f:(Board.of_string)
  |> List.map ~f:value_of
  |> List.map ~f:to_value
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"

let value_of_2 board =
  let (skip_x, skip_y) = match value_of board with
    | (`x, x) -> (x, -1)
    | (`y, y) -> (-1, y)
  in
  let (dimx, dimy) = Board.shape board in
  List.cartesian_product (range dimx) (range dimy)
  |> List.find_map_exn ~f:(fun (x, y) ->
      Board.swap board x y;
      let result =
        try
          Some (value_of ~skip_x ~skip_y board)
        with Failure _ -> 
          None
      in
      Board.swap board x y;
      result)

let part2 content =
  Str.split (Str.regexp "\n\n") content
  |> List.map ~f:(Board.of_string)
  |> List.map ~f:value_of_2
  |> List.map ~f:to_value
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"
