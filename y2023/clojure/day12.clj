(require '[clojure.string :as str])

(defn parse-line [l times]
  (let [str->number-list (fn [s] (map #(Integer/parseInt %) (str/split s #",")))
        [condition damage] (str/split l #" " 2)
        condition (str/join "?" (repeat times condition))
        damage    (str/join "," (repeat times damage))]
    [condition (str->number-list damage)]))

(defn possible?
  "checks wether the collection can start with number damaged entries from the start."
  [coll number]
  (when (>= (count coll) number)
    (and (not (some #(= % \.) (take number coll)))
         (not (= (nth coll number \.) \#)))))

(def m-f
  (memoize
   (fn [condition damage]
     (cond
       (empty? damage)          (if (some #(= % \#) condition) 0 1)
       (empty? condition)       0
       (= \. (first condition)) (m-f (rest condition) damage)
       :else
       (let [number   (first damage)
             possible (possible? condition number)]
         (+ (if possible
              (m-f (drop (inc number) condition) (rest damage)) 0)
            (if (= \? (first condition))
              (m-f (rest condition) damage) 0)))))))

(defn solve [content times]
  (->> content
       str/split-lines
       (pmap #(parse-line % times))
       (pmap (fn [[condition damage]]
               (m-f condition damage)))
       (reduce + 0)))

(comment
  (def content (slurp "../../inputs/2023_12.txt"))
  ;; part1
  (time (solve content 1))
  ;; part2
  (time (solve content 5)))
