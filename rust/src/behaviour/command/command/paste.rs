use super::Command;

#[derive(Default)]
pub struct PasteCommand;

impl Command for PasteCommand {
    fn execute(&mut self, app: &mut std::collections::HashMap<String, String>) {
        println!("Executed paste for {:?}", app);
    }

    fn undo(&mut self, app: &mut std::collections::HashMap<String, String>) {
        println!("Executed undo of paste for {:?}", app);
    }
}
