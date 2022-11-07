main :: IO ();
main = do {
    tinySleep 2000;

    x <- return 5;
    y <- return "some string";
    putStr "Hello!\n";
    putStr "x: ";
    putInt x;
    putStr "\n";
    putStr "y: ";
    putStr y;
    putStr "\nAnother number: ";
    putInt 0x2B;
    putStr "\n";
};