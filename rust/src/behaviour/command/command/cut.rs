use super::Command;

#[derive(Default)]
pub struct CutCommand;

impl Command for CutCommand {
    fn execute(&mut self, app: &mut std::collections::HashMap<String, String>) {
        println!("Executed cut for {:?}", app);
    }

    fn undo(&mut self, app: &mut std::collections::HashMap<String, String>) {
        println!("Executed undo of cut for {:?}", app);
    }
}
