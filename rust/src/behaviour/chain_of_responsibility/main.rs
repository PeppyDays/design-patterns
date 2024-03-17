use crate::behaviour::chain_of_responsibility::{
    department::{
        cashier::Cashier, doctor::Doctor, medical::Medical, reception::Reception, Department,
    },
    patient::Patient,
};

fn main() {
    let cashier = Cashier::default();
    let medical = Medical::new(cashier);
    let doctor = Doctor::new(medical);
    let mut reception = Reception::new(doctor);

    let mut patient = Patient {
        name: "John".into(),
        ..Patient::default()
    };

    reception.execute(&mut patient);

    println!("\nThe patient has been already handled:\n");

    reception.execute(&mut patient);
}
