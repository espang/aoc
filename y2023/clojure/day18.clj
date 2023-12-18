(require '[clojure.string :as str])

(defn parse-line [l]
  (let [[_ direction steps colour] (re-matches #"(\w) (\d*) \((#[0-9a-f]*)\)" l)]
    [direction
     (Integer/parseInt steps)
     colour]))

(defn parse-colour [l]
  [(case (nth l 6)
     \0 "R"
     \1 "D"
     \2 "L"
     \3 "U")
   (Integer/parseInt (subs l 1 6) 16)])

(defn parse [parse-line content]
  (->> (str/split-lines content)
       (map parse-line)))

(defn shoelace'ish [insts]
  (loop [row    0
         col    0
         area   1
         insts' insts]
    (if-not (seq insts')
      area
      (let [[dir steps] (first insts')
            add-area (case dir
                       "R" steps
                       "D" (* steps (inc col))
                       "U" (* -1 steps col)
                       "L" 0)
            row' (case dir
                   "U" (- row steps)
                   "D" (+ row steps)
                   row)
            col' (case dir
                   "L" (- col steps)
                   "R" (+ col steps)
                   col)]
        (recur row'
               col'
               (+ area add-area)
               (rest insts'))))))

(comment
  (def content (slurp "../../inputs/2023_18.txt"))

  ;; part1
  (def insts (parse (fn [l] (take 2 (parse-line l)))
                    content))
  (shoelace'ish insts)

  ;; part2
  (def insts2 (parse (fn [l] (parse-colour (nth (parse-line l) 2)))
                     content))
  (shoelace'ish insts2))
