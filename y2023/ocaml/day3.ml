open Core

module Board = struct
  type t = char array array
  let value_at t row col =
    try t.(row).(col) with Invalid_argument _ -> '.'
end

let is_digit = function | '0'..'9' -> true | _ -> false
let is_no_digit c = not (is_digit c)
let is_symbol = function | '0'..'9' | '.' -> false | _ -> true
let neighbours (row_idx, col_idx) = [
  (row_idx - 1, col_idx - 1);
  (row_idx - 1, col_idx);
  (row_idx - 1, col_idx + 1);
  (row_idx, col_idx - 1);
  (row_idx, col_idx + 1);
  (row_idx + 1, col_idx - 1);
  (row_idx + 1, col_idx);
  (row_idx + 1, col_idx + 1)
]
let parse_board s =
  let rows = String.split_lines s in
  let dimx = List.length rows in
  let dimy = String.length (List.hd_exn rows) in
  let arr = Array.make_matrix ~dimx ~dimy '.' in
  let insert row_idx col_idx cell =
    arr.(row_idx).(col_idx) <- cell
  in
  List.iteri rows ~f:(fun row_idx row ->
    String.iteri row ~f:(insert row_idx));
  arr

let part_numbers board row_idx row =
  let is_symbol_at board (row_idx, col_idx) =
    is_symbol (Board.value_at board row_idx col_idx)
  in
  let has_adjacent_symbol board (position, _) =
    neighbours position
    |> List.exists ~f:(is_symbol_at board)
  in
  let to_part_number lst =
    if List.exists lst ~f:(has_adjacent_symbol board)
    then
      List.map lst ~f:(fun (_, c) -> c)
      |> String.of_list
      |> int_of_string
    else 0
  in
  let char_positions =
    String.to_list row
    |> List.mapi ~f:(fun col_idx c -> ((row_idx, col_idx), c))
  in
  let drop_number lst = List.drop_while lst ~f:(fun (_, c) -> is_digit c) in
  let drop_non_number lst = List.drop_while lst ~f:(fun (_, c) -> is_no_digit c) in
  let rec collect_numbers numbers = function
    | [] -> List.rev numbers
    | lst ->
      let number =
        List.take_while lst ~f:(fun (_, c) -> is_digit c)
        |> to_part_number
      in
      collect_numbers (number::numbers) (drop_number lst |> drop_non_number)
  in
  collect_numbers [] (drop_non_number char_positions)

let part1 content =
  let board = parse_board content in
  String.split_lines content
  |> List.mapi ~f:(part_numbers board)
  |> List.concat
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d"

let gear_numbers board row_idx row =
  let is_gear_symbol_at board (row_idx, col_idx) =
    phys_equal '*' (Board.value_at board row_idx col_idx)
  in
  let gear_symbol board (position, _) =
    neighbours position
    |> List.find ~f:(is_gear_symbol_at board)
  in
  let to_part_number lst =
    let number =
      List.map lst ~f:(fun (_, c) -> c)
      |> String.of_list
      |> int_of_string
    in
    match List.filter_map lst ~f:(gear_symbol board) with
    | [] -> None
    | (row_idx, col_idx) :: _ -> Some ((row_idx, col_idx), number)
  in
  let char_positions =
    String.to_list row
    |> List.mapi ~f:(fun col_idx c -> ((row_idx, col_idx), c))
  in
  let drop_number lst = List.drop_while lst ~f:(fun (_, c) -> is_digit c) in
  let drop_non_number lst = List.drop_while lst ~f:(fun (_, c) -> is_no_digit c) in
  let rec collect_numbers numbers = function
    | [] -> List.rev numbers
    | lst ->
      let number =
        List.take_while lst ~f:(fun (_, c) -> is_digit c)
        |> to_part_number
      in
      collect_numbers (number::numbers) (drop_number lst |> drop_non_number)
  in
  collect_numbers [] (drop_non_number char_positions)

let reduce_gears lst =
  let rec reduce_gears' result = function
    | [] -> result
    | ((row_idx, col_idx), value) :: tl ->
      match List.find tl ~f:(fun ((r, c), _) -> row_idx = r && col_idx = c) with
      | None -> reduce_gears' result tl
      | Some (_, value2) ->
        reduce_gears' (result + value * value2) tl
  in
  reduce_gears' 0 lst

let part2 content =
  let board = parse_board content in
  String.split_lines content
  |> List.mapi ~f:(gear_numbers board)
  |> List.concat
  |> List.filter_map ~f:(fun x -> x)
  |> reduce_gears
  |> Printf.printf "%d"