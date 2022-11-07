main :: IO ();
main = do {
    x <- return 5;
    do {
        x <- return 7;
        putInt x;
        putStr "\n";
    };
    putInt x;
    putStr "\n";
};