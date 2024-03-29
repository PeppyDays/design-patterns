use super::patient::Patient;

pub mod cashier;
pub mod doctor;
pub mod medical;
pub mod reception;

pub trait Department {
    fn handle(&mut self, patient: &mut Patient);
    fn next(&mut self) -> &mut Option<Box<dyn Department>>;
    fn execute(&mut self, patient: &mut Patient) {
        self.handle(patient);

        if let Some(next) = &mut self.next() {
            next.execute(patient);
        }
    }
}

fn into_next(department: impl Department + Sized + 'static) -> Option<Box<dyn Department>> {
    Some(Box::new(department))
}
