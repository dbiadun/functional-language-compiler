main :: IO ();
main = do {
    putInt (fun 2 4);
};

fun :: Int -> Int;
fun x = x + 1;
