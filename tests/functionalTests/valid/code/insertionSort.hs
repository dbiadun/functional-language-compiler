import "List.hs";
import "../../../libs/SortType.hs";

main :: IO ();
main = do {
    l <- return (getList 256);
    l <- return (insertionSort l);

    printList l;
    putStr "\n";
};

insertionSort :: List SortType -> List SortType;
insertionSort l = foldl insert Nil l;

insert :: List SortType -> SortType -> List SortType;
insert l x = case l of {
    Nil -> Cons x Nil;
    Cons y ys -> case cmp x y of {
        True -> Cons x l;
        False -> Cons y (insert ys x);
    };
};
