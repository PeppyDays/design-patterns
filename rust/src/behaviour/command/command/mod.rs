use std::collections::HashMap;

pub mod copy;
pub mod cut;
pub mod paste;

pub trait Command {
    fn execute(&mut self, app: &mut HashMap<String, String>);
    fn undo(&mut self, app: &mut HashMap<String, String>);
}
