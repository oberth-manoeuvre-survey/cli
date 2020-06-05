using System;
using System.IO;
using System.Net;
using Microsoft.Deployment.WindowsInstaller;

namespace InstallStateTool
{
    public class CustomActions
    {
        [CustomAction]
        public static ActionResult InstallStateTool(Session session)
        {
            session.Log("Installing State Tool if necessary");
            if (session.CustomActionData["STATE_TOOL_INSTALLED"] == "true")
            {
                session.Log("State Tool is installed, no installation required");
                return ActionResult.Success;
            }

            string tempDir = Path.GetTempPath();
            string scriptPath = Path.Combine(tempDir, "install.ps1");

            StatusMessage(session, "Installing State Tool...");

            try
            {
                WebClient client = new WebClient();
                client.DownloadFile("https://platform.activestate.com/dl/cli/install.ps1", scriptPath);
            } catch (WebException e)
            {
                session.Log(string.Format("Encoutered exception downloading file: {0}", e.ToString()));
                return ActionResult.Failure;
            }


            string installCmd = string.Format("powershell \"{0} -n\"", scriptPath);
            session.Log(string.Format("Running install command: {0}", installCmd));
            return RunCommand(session, installCmd);
        }

        private static ActionResult RunCommand(Session session, string cmd)
        {
            try
            {
                System.Diagnostics.ProcessStartInfo procStartInfo =
                    new System.Diagnostics.ProcessStartInfo("cmd", "/c " + cmd);

                // The following commands are needed to redirect the standard output.
                // This means that it will be redirected to the Process.StandardOutput StreamReader.
                procStartInfo.RedirectStandardOutput = true;
                procStartInfo.RedirectStandardError = true;
                procStartInfo.UseShellExecute = false;
                // Do not create the black window.
                procStartInfo.CreateNoWindow = true;

                System.Diagnostics.Process proc = new System.Diagnostics.Process();
                proc.StartInfo = procStartInfo;
                proc.Start();
                session.Log(string.Format("Standard output: {0}", proc.StandardOutput.ReadToEnd()));
                session.Log(string.Format("Standard error: {0}", proc.StandardError.ReadToEnd()));
            }
            catch (Exception objException)
            {
                session.Log(string.Format("Caught exception: {0}", objException));
                return ActionResult.Failure;
            }

            return ActionResult.Success;
        }

        internal static void StatusMessage(Session session, string status)
        {
            Record record = new Record(3);
            record[1] = "callAddProgressInfo";
            record[2] = status;
            record[3] = "Incrementing tick [1] of [2]";

            session.Message(InstallMessage.ActionStart, record);
        }

    }
}