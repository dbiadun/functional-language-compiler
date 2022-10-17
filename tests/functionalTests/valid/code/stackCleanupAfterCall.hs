main :: IO ();
main = do {
    a <- return 15;
    putInt a;
    putStr "\n";
    x <- fun;
    putInt a;
    putStr "\n";
    putInt x;
    putStr "\n";
};

fun :: IO Int;
fun = do {
    putStr "Fun start\n";
    a <- return 5;
    putInt a;
    putStr "\n";
    b <- return "new string\n";
    putStr b;
    c <- return 7;
    putStr "Fun end\n";
    return c;
}