public class HelloGoodbye {
    public static void main(String[] args){
        System.out.println("Hello ".concat(args[0]).concat(" and ").concat(args[1]).concat("."));
        System.out.println("Goodbye ".concat(args[1]).concat(" and ").concat(args[0]).concat("."));
    }
}
