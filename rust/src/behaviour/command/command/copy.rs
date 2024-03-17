use super::Command;

#[derive(Default)]
pub struct CopyCommand;

impl Command for CopyCommand {
    fn execute(&mut self, app: &mut std::collections::HashMap<String, String>) {
        println!("Executed copy for {:?}", app);
    }

    fn undo(&mut self, app: &mut std::collections::HashMap<String, String>) {
        println!("Executed undo of copy for {:?}", app);
    }
}
