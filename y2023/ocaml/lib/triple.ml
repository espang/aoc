open Core

module Triple = struct
  module T = struct
    type t = int * int * int
    let compare (x0, y0, z0) (x1, y1, z1) =
      match compare x0 x1 with
      | 0 ->
        (match compare y0 y1 with
        | 0 -> compare z0 z1
        | n -> n)
      | n -> n

    let sexp_of_t = Tuple3.sexp_of_t Int.sexp_of_t Int.sexp_of_t Int.sexp_of_t
    let t_of_sexp = Tuple3.t_of_sexp Int.t_of_sexp Int.t_of_sexp Int.t_of_sexp
  end

  include T
  include Comparable.Make(T)
end
