main :: IO ();
main = do {
    x <- return 5;
    putInt x;
    putStr "\n";
    x <- return (x + 10);
    putInt x;
    putStr "\n";
};