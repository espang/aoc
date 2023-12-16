open Core

let transform current_value c =
  ((current_value + int_of_char c) * 17) % 256

let hash s =
  String.fold s ~init:0 ~f:transform

let part1 content =
  String.split content ~on:','
  |> List.map ~f:hash
  |> List.fold ~init:0 ~f:(+)
  |> Printf.printf "%d\n"

type operation =
  | Delete of int * string
  | Replace of int * (string * int)

let parse_entry s =
  match String.split_on_chars s ~on:['-';'='] with
  | [v; ""] -> Delete ((hash v), v)
  | [v;  n] -> Replace ((hash v), (v, int_of_string n))
  | _ -> failwith "unexpected entry"

let delete boxes key value =
  let rec remove acc = function
    | [] -> List.rev acc
    | (inner_key, _) as hd :: tl ->
      if String.equal value inner_key
      then
        List.append (List.rev acc) tl
      else
        remove (hd::acc) tl
  in
  boxes.(key) <- remove [] boxes.(key)

let add boxes key (inner_key, value) =
  let rec aux acc = function
    | [] -> List.rev ((inner_key, value)::acc)
    | (inner_key', _) as hd :: tl ->
      if String.equal inner_key inner_key'
      then
        List.append (List.rev ((inner_key, value)::acc)) tl
      else
        aux (hd::acc) tl
  in
  boxes.(key) <- aux [] boxes.(key)

let handle boxes = function
  | Delete (key, value) -> delete boxes key value  
  | Replace (key, value) -> add boxes key value

let score boxes =
  let score_list lst multiplier =
    List.foldi lst ~init:0 ~f:(fun i acc (_, v) -> acc + multiplier * (succ i) * v)
  in
  Array.foldi boxes ~init:0 ~f:(fun i acc lst ->
    acc + score_list lst (succ i))

let part2 content =
  let boxes = Array.init 256 ~f:(fun _ -> []) in
  String.split content ~on:','
  |> List.map ~f:parse_entry
  |> List.iter ~f:(handle boxes);
  score boxes
  |> Printf.printf "%d\n"
  
