main :: IO ();
main = do {
    x <- return 15;
    do {
        putInt x;
    };
};