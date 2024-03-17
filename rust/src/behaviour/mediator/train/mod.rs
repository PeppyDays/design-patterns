use super::station::mediator::Mediator;

pub mod freight_train;
pub mod passenger_train;

pub trait Train {
    fn name(&self) -> &str;
    fn arrive(&mut self, mediator: &mut dyn Mediator);
    fn depart(&mut self, mediator: &mut dyn Mediator);
}
