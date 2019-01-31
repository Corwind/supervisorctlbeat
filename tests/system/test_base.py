from supervisorctlbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Supervisorctlbeat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        supervisorctlbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("supervisorctlbeat is running"))
        exit_code = supervisorctlbeat_proc.kill_and_wait()
        assert exit_code == 0
