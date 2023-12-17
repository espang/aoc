open Core

module Pair = struct
  module T = struct
    type t = int * int
    let compare (x0, y0) (x1, y1) =
      match compare x0 x1 with
      | 0 -> compare y0 y1
      | n -> n

    let sexp_of_t = Tuple2.sexp_of_t Int.sexp_of_t Int.sexp_of_t
    let t_of_sexp = Tuple2.t_of_sexp Int.t_of_sexp Int.t_of_sexp
  end

  include T
  include Comparable.Make(T)
end
