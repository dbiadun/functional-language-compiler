main :: IO ();
main = testTermometer;

testTermometer :: IO ();
testTermometer = do {
    tinySleep 2000;

    termometer <- return 0x44;
    i2c <- return 0;

    tinyI2CConfigureDefault i2c;

    w <- return (Cons 0x24 (Cons 0x00 Nil));
    tinyI2CTx i2c termometer w 0;

    tinySetTimer 0 100000 (readTemp i2c termometer);

    loop (putStr "");
};

readTemp :: Int -> Int -> IO ();
readTemp i2c termometer = do {
    res <- tinyI2CTx i2c termometer Nil 6;
    case res of {
        Left err -> putStr "";
        Right l -> do {
            tinyStopTimer 0;

            tempH <- return (getNth 0 l);
            tempL <- return (getNth 1 l);
            humH <- return (getNth 3 l);
            humL <- return (getNth 4 l);

            tempOutput <- return (256 * tempH + tempL);
            humOutput <- return (256 * humH + humL);

            temp <- return (10 * (0 - 45) + 10 * 175 * tempOutput / (256 * 256 - 1));
            tempInt <- return (temp / 10);
            tempRest <- return (temp - 10 * tempInt);
            hum <- return (100 * humOutput / (256 * 256 - 1));

            case tempInt > 10 && tempInt < 40 of {
                True -> putStr "Temperature between 10 and 40\n";
                False -> putStr "Temperature suspiciously low or high\n";
            };
        };
    };
};

loop :: IO a -> IO a;
loop action = do {
    action;
    loop action;
};

getNth :: Int -> List a -> a;
getNth n l = case n > 0 of {
    True -> case l of {
        Cons x xs -> getNth (n - 1) xs;
    };
    False -> case l of {
        Cons x xs -> x;
    };
};
