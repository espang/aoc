open Core

type color =
  | Red
  | Green
  | Blue

let parse_color = function
  | "green" -> Green
  | "red"   -> Red
  | "blue"  -> Blue
  | color -> failwith (Printf.sprintf "unexpected color: '%s'" color)

(* parses a string like 'game 1' and returns the integer id *)
let parse_game s =
  match String.split_on_chars s ~on:[' '] with
  | _ :: game_id :: [] -> int_of_string game_id
  | _ -> failwith (Printf.sprintf "unexpected game pattern: '%s'" s)

(* parse_cube taks a string like '6 blue' a tuple with the color and number *)
let parse_cube s =
  match String.split_on_chars s ~on:[' '] with
  | number :: color :: [] -> (parse_color color, int_of_string number)
  | _ -> failwith (Printf.sprintf "unexpected cube pattern: '%s'" s)

(* parse_set parses a comma separated list of cubes *)
let parse_set s =
  String.split_on_chars s ~on:[',']
  |> List.map ~f:String.strip
  |> List.map ~f:parse_cube

(* parse_set parses a semicolon separated list of sets *)
let parse_sets s =
  String.split_on_chars s ~on:[';']
  |> List.map ~f:String.strip
  |> List.map ~f:parse_set

let parse_line l =
  match String.split_on_chars l ~on:[':'] |> List.map ~f:String.strip with
  | game_pattern :: set_patterns :: [] ->
    (parse_game game_pattern, parse_sets set_patterns)
  | _ -> failwith (Printf.sprintf "unexpected line: '%s'" l)

let game_possible sets (setup:(color * int) list) =
  let rec check_set = function
    | [] -> true
    | (color, number) :: tl ->
      let max_number = List.Assoc.find_exn setup ~equal:phys_equal color in
      if number > max_number then false else check_set tl
  in
  let rec gp = function
    | [] -> true
    | set :: tl ->
      (if check_set set then gp tl else false)
  in
  gp sets

let score_line setup l =
  let (game_id, sets) = parse_line l in
  if game_possible sets setup
  then game_id
  else 0

let part1 content =
  let setup = [(Red, 12); (Green, 13); (Blue, 14)] in
  String.split_lines content
  |> List.map ~f:(score_line setup)
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d"

let value_for color set =
  List.Assoc.find set ~equal:phys_equal color
  |> Option.value ~default:1

let score_line_2 l =
  let (_, sets) = parse_line l in
  [Green; Red; Blue]
  |> List.map ~f:(fun color ->
    List.map sets ~f:(value_for color)
    |> List.max_elt ~compare
    |> Option.value_exn)
  |> List.fold ~init:1 ~f:( * )

let part2 content =
  String.split_lines content
  |> List.map ~f:score_line_2
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d"
