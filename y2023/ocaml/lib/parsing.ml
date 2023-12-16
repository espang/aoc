open Core

let range i = List.init i ~f:(fun x -> x)
let range_rev i = List.init i ~f:(fun x -> i - x -1)

let parse_string_matrix char_to_t s =
  let rows =
    String.split_lines s
    |> List.map ~f:String.to_list
    |> List.map ~f:(List.map ~f:char_to_t)
  in
  let nrows = List.length (List.hd_exn rows) in
  let ncols = List.length rows in
  List.map (range nrows) ~f:(fun row ->
    Array.init ncols ~f:(fun col ->
      List.nth_exn (List.nth_exn rows row) col))
  |> Array.of_list

type shape = {nrows:int; ncols:int}
let shape_of matrix =
  {nrows=Array.length matrix; ncols=Array.length matrix.(0)}