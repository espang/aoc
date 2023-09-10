(*

$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k

*)

(* command: start with '$ '*)

type cmd =
  | CD of string
  | LS
  | EOF

type entry =
  | File of string * int
  | Dir of string * entry list

let read_file_system s =
  let rec next_command lines =
    match lines with
    | [] -> (EOF, [])
    | hd :: tl ->
      let splitted = String.split_on_char ' ' hd in
      match splitted with
      | [a; b] when a = "$" && b = "ls" -> (LS, tl)
      | [a; b; c] when a = "$" && b = "cd" -> (CD c, tl)
      | _ -> failwith ("invalid command " ^ hd)
  in
  let rec until_next_command [] = function
  
    []
  in
  let rec resolve_fs acc expect_command lines =
    if expect_command
    then
      match next_command lines with
      | (EOF, _) -> acc
      | (LS, new_lines) -> 