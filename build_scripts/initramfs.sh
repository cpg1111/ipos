#!/bin/bash

mkdir -p $MOUNT_PATH/usr/src/initramfs/{bin,dev,etc,lib,lib64,mnt/root,proc,root,sbin,sys}
cp -a /dev/{null,console,tty, $TARGET_PARITION} $MOUNT_PATH/usr/src/initramfs/dev/
USE="static" emerge -av busybox
cp -a /bin/busybox $MOUNT_PATH/usr/src/initramfs/bin/busybox
cp $(dirname $0)/../dist/fuse $MOUNT_PATH/usr/src/initramfs/bin/fuse
cp $(dirname $0)/../dist/ipfs $MOUNT_PATH/usr/src/initramfs/bin/ipfs
cp $(dirname $0)/../boot/init $MOUNT_PATH/usr/src/initramfs/init
cp -r $(dirname $0)/../etc/ $MOUNT_PATH/usr/src/initramfs/etc/
echo "mount -o $TARGET_PARTITION /mnt/root" >> $MOUNT_PATH/usr/src/initramfs/init
echo "mkdir -p /mnt/root/var/ /mnt/root/opt/ /mnt/root/home/ /mnt/root/usr/bin/ /mnt/root/usr/local/ /mnt/root/etc/ /mnt/root/root/ /mnt/root/tmp/ /mnt/root/lib/" >> $MOUNT_PATH/usr/src/initramfs/init
echo "ipfs config --json Mounts.FuseAllowOther true" >> $MOUNT_PATH/usr/src/initramfs/init
echo "ipfs daemon --mount" >> $MOUNT_PATH/usr/src/initramfs/init
