pragma circom 2.0.0;

/*This circuit template checks that c is the multiplication of a and b.*/  
//tengo que hacer un parametro de entrada vector (un array) que coja todos los numero y los sume por ejemplo. Entonces en el contrato ir haciendo la suma de todos los eventos  y comparar los resultados.
//Basicamente lo que hago aqui es probar que he sincronizado todos los eventos porque la suma es la correcta. No tengo en cuenta la propiedad conmutativa de la suma.
template Sum () {  

   // Declaration of signals.  
   signal input a;  
   signal input b;  
   signal output c;  

   // Constraints.  
   c <== a + b;  
}

template Hackaton (N){
   //Declaration of signals and components.
   signal input in[N];
   //signal input2 in[256];
   signal output out;
   component comp[N-1];

   //Statements.
   for(var i = 0; i < N-1; i++){
       comp[i] = Sum();
   }
}

component main = Hackaton(1000);
