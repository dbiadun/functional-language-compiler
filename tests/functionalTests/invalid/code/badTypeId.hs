data bool = T | F;

main :: IO ();
main = case T of {
    T -> putStr "True";
    F -> putStr "False";
};