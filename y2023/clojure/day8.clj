#!/usr/bin/env bb
(require
 '[clojure.math.numeric-tower :as math]
 '[clojure.string :as str])

(defn parse [network]
  (let [parse-line (fn [l]
                     (let [pattern #"(\w\w\w) = \((\w\w\w), (\w\w\w)\)"
                           [_ start left right] (re-matches pattern l)]
                       [start [left right]]))]
    (->> (str/split-lines network)
         (map parse-line)
         (into {}))))

(defn parse-input [s]
  (let [[instructions network] (str/split s #"\n\n")]
    [instructions
     (parse network)]))

(defn move [network instruction position]
  (if-let [[left right] (get network position nil)]
    (case instruction
      \R right
      \L left)
    (throw (Exception. "unknown position"))))

(defn walk [from to instructions network]
  (loop [steps         0
         instructions' (cycle (seq instructions))
         position      from]
    (if (= position to)
      steps
      (recur (inc steps)
             (rest instructions')
             (move network (first instructions') position)))))

(defn cycle-len [from inst-len instructions network]
  (loop [steps         0
         instructions' instructions
         position      from]
    (if (and (str/ends-with? position "Z")
             (zero? (mod steps inst-len)))
      steps
      (recur (inc steps)
             (rest instructions')
             (move network (first instructions') position)))))

(defn solve-part2 [instructions network]
  (let [inst-len (count instructions)]
    (->> (keys network)
         (filter #(str/ends-with? % "A"))
         (map #(cycle-len %
                          inst-len
                          (cycle (seq instructions))
                          network))
         (reduce (fn [acc v] (math/lcm acc v))
                 1))))

(comment
  (def content (slurp "../../inputs/2023_8.txt"))
  (defn part1 [content]
    (let [[instructions network] (parse-input content)]
      (walk "AAA" "ZZZ" instructions network)))
  (defn part2 [content]
    (let [[instructions network] (parse-input content)]
      (solve-part2 instructions network)))
  (part1 content)
  (part2 content))
