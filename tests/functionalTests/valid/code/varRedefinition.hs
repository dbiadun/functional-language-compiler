main :: IO ();
main = do {
    x <- return 5;
    putInt x;
    putStr "\n";
    x <- return 8;
    putInt x;
    putStr "\n";
};