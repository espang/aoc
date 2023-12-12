(require '[clojure.string :as str])

(defn str->number-list [s] (map #(Integer/parseInt %) (str/split s #",")))

(defn parse-line [l times]
  (let [[condition damage] (str/split l #" " 2)
        condition (str/join "?" (repeat times condition))
        damage    (str/join "," (repeat times damage))]
    [condition (str->number-list damage)]))

(defn is-operational? [c] (= c \.))
(defn is-damaged? [c] (= c \#))

(defn can-be?
  [coll number]
  (when (>= (count coll) number)
    (and (not (some is-operational? (take number coll)))
         (not (= (nth coll number \.) \#)))))

(def m-f
  (memoize
   (fn [condition damage]
     (if-not (seq damage)
       (if (some is-damaged? condition)
         0
         1)
       (if-not (seq condition)
         0
         (let [number (first damage)]
           (case (first condition)
             \# (if (can-be? condition number)
                  (m-f (drop (inc number) condition) (rest damage))
                  0)
             \? (if (can-be? condition number)
                  (+ (m-f (drop (inc number) condition) (rest damage))
                     (m-f (rest condition) damage))
                  (m-f (rest condition) damage))
             \. (m-f (rest condition) damage))))))))

(comment
  (def content (slurp "../../inputs/2023_12.txt"))
  ;; part1
  (time (->> content
             str/split-lines
             (map #(parse-line % 1))
             (map (fn [[condition damage]]
                    (m-f condition damage)))
             (reduce + 0)))
  (time (->> content
             str/split-lines
             (map #(parse-line % 5))
             (map (fn [[condition damage]]
                    (m-f condition damage)))
             (reduce + 0))))
