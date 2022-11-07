main :: IO ();
main = do {
    tinySleep 2000;

    tinySetTimer 0 2000000 timer0Handler;
    tinySetTimer 1 500000 timer1Handler;
    tinySetTimer 2 50000 timer2Handler;
    tinySetTimer 3 100000 timer3Handler;
    tinyStopTimer 3;

    loop (putStr "");
};

timer0Handler :: IO ();
timer0Handler = do {
    putStr "Timer ";
    putInt 0;
    putStr "\n";
    tinyStopTimer 0;
};

timer1Handler :: IO ();
timer1Handler = do {
    n <- stateGetInt "timer1Counter"; -- 0 is the default value
    case n < 4 of {
        True -> do {
            stateSetInt "timer1Counter" (n + 1);

            putStr "Timer ";
            putInt 1;
            putStr "\n";
        };
        False -> tinyStopTimer 1;
    };

    case n == 1 of {
        True -> tinyStartTimer 2;
        False -> putStr "";
    };

    case n == 2 of {
        True -> tinySetTimer 4 30000 timer4Handler;
        False -> putStr "";
    };
};

timer2Handler :: IO ();
timer2Handler = do {
    n <- stateGetInt "timer2Counter"; -- 0 is the default value
    case n < 5 of {
        True -> do {
            stateSetInt "timer2Counter" (n + 1);

            putStr "Timer ";
            putInt 2;
            putStr "\n";
        };
        False -> tinyStopTimer 2;
    };

    case n == 2 of {
        True -> do {
            tinyStopTimer 2;
            tinyStartTimer 3;
        };
        False -> putStr "";
    };
};

timer3Handler :: IO ();
timer3Handler = do {
    n <- stateGetInt "timer3Counter"; -- 0 is the default value
    case n < 2 of {
        True -> do {
            stateSetInt "timer3Counter" (n + 1);

            putStr "Timer ";
            putInt 3;
            putStr "\n";
        };
        False -> tinyStopTimer 3;
    };
};

timer4Handler :: IO ();
timer4Handler = do {
    n <- stateGetInt "timer4Counter"; -- 0 is the default value
    case n < 6 of {
        True -> do {
            stateSetInt "timer4Counter" (n + 1);

            putStr "Timer ";
            putInt 4;
            putStr "\n";
        };
        False -> tinyStopTimer 4;
    };
};

loop :: IO a -> IO a;
loop action = do {
    action;
    loop action;
};
