open Core

type space = Empty | Galaxy

let space_of_char = function
  | '.' -> Empty
  | '#' -> Galaxy
  | c -> failwith (Printf.sprintf "unexpected char '%c'" c)

let parse content =
  let lst_matrix = String.split_lines content
    |> List.map ~f:String.to_list
  in
  let dimx = List.length (List.hd_exn lst_matrix) in
  let dimy = List.length lst_matrix in
  Array.init dimx ~f:(fun column ->
    Array.init dimy ~f:(fun row ->
      List.nth_exn (List.nth_exn lst_matrix row) column
      |> space_of_char))

let empty_x board =
  Array.filter_mapi board ~f:(fun x arr ->
    if Array.for_all arr ~f:(phys_equal Empty)
    then Some x
    else None)
  |> Array.to_list

let empty_y board =
  let range i = List.init i ~f:(fun x -> x) in
  let dimx = Array.length board in
  let dimy = Array.length board.(0) in
  range dimy
  |> List.filter ~f:(fun y ->
    range dimx
    |> List.map ~f:(fun x -> board.(x).(y))
    |> List.for_all ~f:(phys_equal Empty))

let transform ?(multiplier = 2) empty_xs empty_ys (x, y) =
  let x_count = List.count empty_xs ~f:(fun x' -> x' < x) in
  let y_count = List.count empty_ys ~f:(fun y' -> y' < y) in
  (x + x_count * (multiplier - 1),
   y + y_count * (multiplier - 1))

let distances ?(multiplier = 2) image =
  let empty_xs = empty_x image in
  let empty_ys = empty_y image in
  let manhattan_distance (x1, y1) (x2, y2) =
    Int.abs (x2-x1) + Int.abs(y2-y1)
  in
  let rec total_distance total = function
    | []     -> total
    | hd::tl ->
      let sub_total =
        List.map tl ~f:(manhattan_distance hd)
        |> List.fold ~init:0 ~f:(+)
      in
      total_distance (total + sub_total) tl
  in
  Array.concat_mapi image ~f:(fun x arr ->
    Array.filter_mapi arr ~f:(fun y c ->
      if phys_equal c Galaxy
      then Some (transform ~multiplier empty_xs empty_ys (x, y))
      else None))
  |> Array.to_list
  |> total_distance 0

let part1 content = 
  parse content
  |> distances
  |> Printf.printf "%d\n"

let part2 content = 
  parse content
  |> distances ~multiplier:1000000
  |> Printf.printf "%d\n"
